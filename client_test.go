package client

import (
	"net/http"
	"testing"
)

func TestTrace(t *testing.T) {
	if req, err := http.NewRequest(http.MethodGet, "https://ident.me/", nil); err != nil {
		t.Errorf("TestTraceError: %v\n", err)

	} else {
		var client = NewClient(NewTransport(nil))
		if err := SetCookieHandler(client); err != nil {
			t.Errorf("TestTraceError: %v\n", err)
		}

		req = SetRequestTrace(req, nil)
		if resp, err := client.Do(req); err != nil {
			t.Errorf("TestTraceError: %v\n", err)

		} else {
			defer resp.Body.Close()
			//io.Copy(os.Stdout, resp.Body)
		}
	}
}
