package usecase

import (
	"context"
	"crypto/rand"
	"math/big"
)

const difficulty = 5

type powUsecase struct {
	ps powStorage
}

func NewPOW(ps powStorage) *powUsecase {
	return &powUsecase{ps: ps}
}

func (u *powUsecase) NewRequest(ctx context.Context) (out string, err error) {
	nonce, err := generateNonce(5)
	if err != nil {
		return
	}
	return u.ps.NewRequest(difficulty, nonce), err
}

func (u *powUsecase) Check(challenge, pow string) (bool, error) {
	return u.ps.Check(challenge, pow)
}

func generateNonce(n int) (out []byte, err error) {
	const symbols = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
	symbolsLen := int64(len(symbols))
	out = make([]byte, n)
	var idx *big.Int
	for i := range out {
		idx, err = rand.Int(rand.Reader, big.NewInt(symbolsLen))
		if err != nil {
			return
		}
		out[i] = symbols[idx.Uint64()]
	}
	return
}
