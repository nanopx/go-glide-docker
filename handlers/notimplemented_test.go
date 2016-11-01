package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestNotImplemented(t *testing.T) {
	ts := httptest.NewServer( NotImplemented )
	defer ts.Close()

	res, err := http.Get( ts.URL )
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}