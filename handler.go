package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Handler struct {
	http.Handler
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	f := strings.Replace(req.RequestURI, "wails://wails", "", 1)
	f, err = url.PathUnescape(f)

	if err != nil {
		fmt.Println("🚨 ", err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not clean up file name %s", f)))
	}
	c, err := os.ReadFile(f)

	if err != nil {
		fmt.Println("🚨 ", err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not serve file %s", f)))
	}

	res.Write(c)
}
