package main

import (
	"context"
	"fmt"
	"net/http"
	"log"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":10200",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to close: %s", err)
			return err
		}
		return nil
	})

	<-ctx.Done()

	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("Failed to shutdown: %s", err)
	}

	return eg.Wait()

}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}

}
