package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	m := routes()

	srv := http.Server{
		Addr:    ":8080",
		Handler: m,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("not able to start server ", err)
		}
	}()

	log.Println("server listening on port", srv.Addr)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sig := <-c

	ctx := context.Background()

	log.Println("Got signal", sig, "shutting server down")
	srv.Shutdown(ctx)
}
