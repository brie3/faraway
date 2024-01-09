package server

import (
	"context"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/brie3/faraway/internal/handler"
	"github.com/brie3/faraway/internal/storage"
	"github.com/brie3/faraway/internal/usecase"
	"github.com/brie3/faraway/pkg/config"
)

type quotes struct {
	server *fiber.App
	log    zerolog.Logger
}

func NewQuotes() *quotes {
	quotesStorage := storage.NewQuotes()
	powStorage := storage.NewPOW()

	quotesUsecase := usecase.NewQuotes(quotesStorage)
	powUsecase := usecase.NewPOW(powStorage)

	out := quotes{
		server: fiber.New(),
		log:    log.Logger.With().Str("module", "server").Logger(),
	}

	handler.Quotes(out.server, quotesUsecase, powUsecase)
	return &out
}

func (a *quotes) MustRun(ctx context.Context, g *errgroup.Group) {
	g.Go(func() error {
		return a.server.Listen(config.MustGet().ServiceBind)
	})
	a.log.Debug().Msg("started")

	<-ctx.Done()

	if err := a.server.ShutdownWithContext(ctx); err != nil {
		a.log.Err(err).Send()
	}
	a.log.Debug().Msg("stopped")
}
