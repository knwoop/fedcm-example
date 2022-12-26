package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/knwoop/fedcm-example/idp/http"
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	s := http.NewServer(8080)
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-termCh:
		return 0
	case err := <-errCh:
		fmt.Fprintf(os.Stderr, "[ERROR] error server running: %s", err)
		return 1
	}
}
