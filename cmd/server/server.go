package main

import (
	"github.com/rs/zerolog/log"

	"github.com/brie3/faraway/internal/app/server"
	"github.com/brie3/faraway/pkg/config"
	"github.com/brie3/faraway/pkg/graceful"
)

func main() {
	log.Logger = config.MustGet().Logger()
	quotesService := server.NewQuotes()

	g, ctx := graceful.New()
	quotesService.MustRun(ctx, g)

	if err := g.Wait(); err != nil {
		log.Logger.Err(err).Send()
	}
}
