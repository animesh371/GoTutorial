package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Serving request from new web server</h1>")
}
