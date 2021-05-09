package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"workplace/src/github.com/pkg/errors"
)

var serverSignal = make(chan struct{})

func main(){
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()

	mux.HandleFunc("/h1", handler1)
	mux.HandleFunc("/h2", handler2)
	mux.HandleFunc("/h3", handler3)

	sever := &http.Server{
		Addr: "127.0.0.1:8000",
		Handler: mux,
	}

	g.Go(func() error {
		return sever.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <- ctx.Done():
			log.Printf("errgroup exit...")
		case <-serverSignal:
			log.Println("server quit...")
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		log.Println("shutting down server...")
		return sever.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("Get os signal: %v", sig)
		}
	})
	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}

func handler1(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Hello handler1...")
}

func handler2(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Hello handler2...")
}

func handler3(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Goodbye handler3...")
	serverSignal <- struct{}{}
}
