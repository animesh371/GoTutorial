package handlers

import (
	"fmt"
	"net/http"
)

var Index int

func init() {
	Index = 5
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Serving request from new web server</h1>")
}
