package main

import (
	"io"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "api2 response")
}
