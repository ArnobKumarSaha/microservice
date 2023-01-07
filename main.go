package main

import (
	"context"
	"flag"
	"github.com/Arnobkumarsaha/microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var bindAddr string

func init()  {
	flag.StringVar(&bindAddr, "BIND_ADDR", ":9090", "usage")
}

func main() {
	lg := log.New(os.Stdout, "my-simple-api", log.LstdFlags)
	mux := http.NewServeMux()

	productHandler := handlers.NewProducts(lg)
	mux.Handle("/", productHandler)


	server := http.Server{
		Addr:              bindAddr,
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


