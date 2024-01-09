package storage

import (
	"context"
	_ "embed"
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/brie3/faraway/pkg/model"
)

//go:embed quotes.json
var quotesString string

type quotes struct {
	log zerolog.Logger
}

func NewQuotes() *quotes {
	return &quotes{
		log: log.Logger.With().Str("module", "quotes").Logger(),
	}
}

func (s *quotes) GetQuotes(_ context.Context) (out []model.Quote, err error) {
	err = json.Unmarshal([]byte(quotesString), &out)
	return
}
