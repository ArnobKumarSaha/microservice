package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	lg *log.Logger
}

func NewGoodbye(lg *log.Logger) *Goodbye {
	return &Goodbye{lg: lg}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.lg.Println("handling goodbye request")
	fmt.Fprintf(rw, "Goodbye")
}