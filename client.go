package client

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptrace"
	"os"

	"golang.org/x/net/publicsuffix"
)

func NewTlsConfig(skipVerify bool) *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: skipVerify,
	}
}

func NewTransport(tlsConfig *tls.Config) *http.Transport {
	return &http.Transport{
		MaxIdleConns:        DefaultMaxIdleConns,
		IdleConnTimeout:     DefaultIdleConnTimeout,
		MaxConnsPerHost:     DefaultMaxConnsPerHost,
		MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
		TLSClientConfig:     tlsConfig,
	}
}

func NewClient(transport *http.Transport) *http.Client {
	return &http.Client{
		Timeout:   DefaultClientTimeout,
		Transport: transport,
	}
}

func NewClientTrace(logger *log.Logger) httptrace.ClientTrace {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}

	return httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			logger.Println("Starting to create connection", hostPort)
		},

		DNSStart: func(info httptrace.DNSStartInfo) {
			logger.Println("Starting to look up dns", info.Host)
		},

		DNSDone: func(info httptrace.DNSDoneInfo) {
			if info.Err == nil {
				logger.Println("Done looking up dns", info.Addrs[0].IP)

			} else {
				logger.Println(info.Err.Error())
			}
		},

		ConnectStart: func(network, addr string) {
			logger.Println("Starting tcp connection to", addr)
		},

		ConnectDone: func(network, addr string, err error) {
			if err == nil {
				logger.Println("Connection created to", addr)

			} else {
				logger.Println(err.Error())
			}
		},

		GotConn: func(info httptrace.GotConnInfo) {
			logger.Println("Connection established", info.Conn.LocalAddr(), "->", info.Conn.RemoteAddr())
		},

		TLSHandshakeStart: func() {
			logger.Print("TLS handshake is started...")
		},

		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			if err == nil {
				logger.Println("Successful handshake connection state to", state.ServerName)
				logger.Println("Peer certificate", state.PeerCertificates[0].Issuer.CommonName, state.PeerCertificates[0].NotAfter)

			} else {
				logger.Println(err.Error())
			}
		},
	}
}

func SetRequestTrace(req *http.Request, logger *log.Logger) *http.Request {
	var trace = NewClientTrace(logger)
	return req.WithContext(httptrace.WithClientTrace(req.Context(), &trace))
}

func SetCookieHandler(client *http.Client) error {
	var options = cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}

	if cookie, err := cookiejar.New(&options); err != nil {
		return err

	} else {
		client.Jar = cookie
		return nil
	}
}

func GetBodyReader(resp *http.Response) *bufio.Reader {
	return bufio.NewReader(resp.Body)
}

func GetJsonDecoder(resp *http.Response) *json.Decoder {
	return json.NewDecoder(resp.Body)
}
