package handlers

import (
	"net/http"
)

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
}

var NotImplemented = http.HandlerFunc(notImplemented)
