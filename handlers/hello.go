package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello")
	name, ok := r.URL.Query()["name"]
	if ok && len(name) > 0 {
		fmt.Fprintf(wr, "Hello, %s", name[0])
	} else {
		fmt.Fprintf(wr, "Hello, World")
	}

}
