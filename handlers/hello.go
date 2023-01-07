package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	lg *log.Logger
}

func NewHello(lg *log.Logger) *Hello {
	return &Hello{lg: lg}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.lg.Println("handling hello request")
	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Can't read from request body", http.StatusBadRequest)
		return
	}
	_, err = fmt.Fprintf(rw, fmt.Sprintf("Hello %s\n", content))
	if err != nil {
		http.Error(rw, "Can't write to the response writer", http.StatusInternalServerError)
		return
	}
}