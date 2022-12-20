package http

import (
	"fmt"
	"net/http"
)

var (
	_ http.Handler = (*helloHandler)(nil)
)

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "hello")
}
