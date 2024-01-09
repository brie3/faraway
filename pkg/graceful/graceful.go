package graceful

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// New returns errgroup with context that terminates on signal
// or error is received in any of the goroutines.
func New() (*errgroup.Group, context.Context) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		defer func() {
			signal.Stop(shutdown)
			close(shutdown)
		}()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-shutdown:
			return errors.New("shutdown")
		}
	})
	return g, ctx
}
