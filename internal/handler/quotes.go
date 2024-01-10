package handler

import (
	"net/http"

	"github.com/brie3/faraway/pkg/model"
	fiber "github.com/gofiber/fiber/v2"
)

type quotesHandler struct {
	quotes quotesUsecase
	pows   powUsecase
}

func Quotes(router *fiber.App, quotes quotesUsecase, pows powUsecase) {
	h := quotesHandler{
		quotes: quotes,
		pows:   pows,
	}
	router.Use(h.checkPOW)
	router.Get("/quote", h.getQuotes)
}

func (h *quotesHandler) getQuotes(ctx *fiber.Ctx) (err error) {
	quote, err := h.quotes.GetQuote(ctx.Context())
	if err != nil {
		return
	}
	return ctx.JSON(quote)
}

func (h *quotesHandler) checkPOW(ctx *fiber.Ctx) (err error) {
	challenge := ctx.Get(model.ChallengeKey)
	pow := ctx.Get(model.POWKey)
	if len(challenge) == 0 {
		challenge, err = h.pows.NewRequest(ctx.Context())
		if err != nil {
			return
		}
		ctx.JSON(model.Response{Challenge: challenge})
		return
	}
	ok, err := h.pows.Check(challenge, pow)
	if err != nil {
		return
	}
	if !ok {
		return ctx.Status(http.StatusBadRequest).SendString("challenge not solved")
	}
	return ctx.Next()
}
