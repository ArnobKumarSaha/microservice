package main

import (
	"github.com/Arnobkumarsaha/microservice/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	lg := log.New(os.Stdout, "my-simple-api", log.LstdFlags)
	mux := http.NewServeMux()

	helloHandler := handlers.NewHello(lg)
	mux.Handle("/", helloHandler)

	goodbyeHandler := handlers.NewGoodbye(lg)
	mux.Handle("/bye", goodbyeHandler)

	http.ListenAndServe(":9090", mux)
}


