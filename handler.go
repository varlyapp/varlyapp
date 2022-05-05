package main

import (
	"fmt"
	"net/http"
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
	f := strings.Replace(req.RequestURI, "wails://wails", "", 1)
	c, err := os.ReadFile(f)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not serve file %s", f)))
	}

	res.Write(c)
}
