package storage

import (
	gopow "github.com/bwesterb/go-pow"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type pow struct {
	log zerolog.Logger
}

func NewPOW() *pow {
	return &pow{
		log: log.Logger.With().Str("module", "pow").Logger(),
	}
}

func (p *pow) NewRequest(difficulty uint32, nonce []byte) string {
	return gopow.NewRequest(difficulty, nonce)
}

func (p *pow) Check(request, pow string) (bool, error) {
	return gopow.Check(request, pow, nil)
}
