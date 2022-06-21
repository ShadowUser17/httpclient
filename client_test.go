package client

import (
	"fmt"
	"net/http"
	"testing"
)

type identMe struct {
	Ip        string `json:"ip"`
	Aso       string `json:"aso"`
	Asn       string `json:"asn"`
	Continent string `json:"continent"`
	Code      string `json:"cc"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Postal    string `json:"postal"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	TimeZone  string `json:"tz"`
}

func (im *identMe) String() string {
	return fmt.Sprint(
		"\nIP:        "+im.Ip,
		"\nASO:       "+im.Aso,
		"\nASN:       "+im.Asn,
		"\nContinent: "+im.Continent,
		"\nCode:      "+im.Code,
		"\nCountry:   "+im.Country,
		"\nCity:      "+im.City,
		"\nPostal:    "+im.Postal,
		"\nLatitude:  "+im.Latitude,
		"\nLongitude: "+im.Longitude,
		"\nTimeZone:  "+im.TimeZone,
	)
}

func TestTrace(t *testing.T) {
	if req, err := http.NewRequest(http.MethodGet, "https://ident.me/", nil); err != nil {
		t.Errorf("NewRequest: %v\n", err)

	} else {
		var client = NewClient(NewTransport(nil))
		if err := SetCookieHandler(client); err != nil {
			t.Errorf("SetCookieHandler: %v\n", err)
		}

		req = SetRequestTrace(req, nil)
		if resp, err := client.Do(req); err != nil {
			t.Errorf("SetRequestTrace: %v\n", err)

		} else {
			defer resp.Body.Close()

			var body = GetBodyReader(resp)
			if line, _, err := body.ReadLine(); err != nil {
				t.Errorf("GetBodyReader: %v\n", err)

			} else {
				t.Logf("Line of body: %s\n", line)
			}
		}
	}
}

func TestJson(t *testing.T) {
	if req, err := http.NewRequest(http.MethodGet, "https://ident.me/json", nil); err != nil {
		t.Errorf("NewRequest: %v\n", err)

	} else {
		var client = NewClient(NewTransport(nil))
		if err := SetCookieHandler(client); err != nil {
			t.Errorf("SetCookieHandler: %v\n", err)
		}

		if resp, err := client.Do(req); err != nil {
			t.Errorf("client.Do: %v\n", err)

		} else {
			defer resp.Body.Close()

			var body = GetJsonDecoder(resp)
			var data = new(identMe)

			if err = body.Decode(data); err != nil {
				t.Errorf("body.Decode: %v\n", err)

			} else {
				t.Logf("JSON of body: %s\n", data)
			}
		}
	}
}
