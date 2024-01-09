package usecase

import (
	"context"

	"github.com/brie3/faraway/pkg/model"
)

type quotesStorage interface {
	GetQuotes(ctx context.Context) (out []model.Quote, err error)
}

type powStorage interface {
	NewRequest(difficulty uint32, nonce []byte) string
	Check(request, pow string) (bool, error)
}
