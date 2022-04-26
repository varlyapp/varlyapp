package main

import (
	"fmt"
	"net/http"
	"os"
)

type Handler struct {
	http.Handler
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	c, err := os.ReadFile(req.RequestURI)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not serve file %s", req.RequestURI)))
	}

	// res.Header().Add("", "")

	res.Write(c)
}
