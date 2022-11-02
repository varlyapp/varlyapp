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
	var err error
	path := req.URL.Path
	f, err := os.ReadFile(path)

	fmt.Println("path", path)
	fmt.Println("req.URL", req.URL.String())

	if err != nil {
		fmt.Println("🚨 ", err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not serve file %s", path)))
	}

	res.Write(f)
}
