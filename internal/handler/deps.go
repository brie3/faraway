package handler

import (
	"context"

	"github.com/brie3/faraway/pkg/model"
)

type quotesUsecase interface {
	GetQuote(ctx context.Context) (out model.Quote, err error)
}

type powUsecase interface {
	NewRequest(ctx context.Context) (out string, err error)
	Check(challenge, pow string) (bool, error)
}
