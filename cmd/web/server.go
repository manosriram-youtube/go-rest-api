package main

import (
	"context"
	handler "go-rest/pkg/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "Logger-REST ", log.LstdFlags)

	// Routes
	hh := handler.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := http.Server{
		Addr:     ":8080",
		Handler:  sm,
		ErrorLog: l,
	}
	go func() {
		l.Println("Server at 8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Create a channel to listen to OS signals.
	c := make(chan os.Signal, 1)

	// Notify the channel on following events -> Interrupt, Kill
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Code blocks here until an Interrupt or Kill signal is received, ie any message is sent to the channel.
	sig := <-c
	l.Println("Got Signal: ", sig)

	// Shutdown server gracefully
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
