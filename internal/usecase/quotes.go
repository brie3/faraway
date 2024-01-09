package usecase

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/brie3/faraway/pkg/model"
)

type quotesUsecase struct {
	qs quotesStorage
}

func NewQuotes(qs quotesStorage) *quotesUsecase {
	return &quotesUsecase{qs: qs}
}

func (u *quotesUsecase) GetQuote(ctx context.Context) (out model.Quote, err error) {
	quotes, err := u.qs.GetQuotes(ctx)
	if err != nil {
		return
	}
	idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(quotes))))
	if err != nil {
		return
	}
	out = quotes[idx.Uint64()]
	return
}
