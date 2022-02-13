package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, errCtx := errgroup.WithContext(context.Background())

	appServer := NewAppServer()

	g.Go(func() error {
		return appServer.ListenAndServe()
	})

	g.Go(func() error {
		<-errCtx.Done()
		log.Println("shutting down server...")
		return appServer.Shutdown(errCtx)
	})

	g.Go(func() error {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-errCtx.Done():
			return errCtx.Err()
		case sig := <-ch:
			return errors.Errorf("receive signal: %v", sig)
		}
	})

	log.Printf("%v\n", g.Wait())
}

func NewAppServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
