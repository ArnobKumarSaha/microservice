package main

import (
	"context"
	"github.com/Arnobkumarsaha/microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	lg := log.New(os.Stdout, "my-simple-api", log.LstdFlags)
	mux := http.NewServeMux()

	helloHandler := handlers.NewHello(lg)
	mux.Handle("/", helloHandler)

	goodbyeHandler := handlers.NewGoodbye(lg)
	mux.Handle("/bye", goodbyeHandler)


	server := http.Server{
		Addr:              ":9090",
		Handler:           mux,
		ReadTimeout:       time.Second * 10,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       time.Minute * 5,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	receivedSignal := <- sigChan
	lg.Printf("Received signal %s for graceful Shutdown", receivedSignal)
	timeout, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	_ = server.Shutdown(timeout)
}


