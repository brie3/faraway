package model

const (
	ChallengeKey string = `X-User-Challenge`
	POWKey       string = `X-User-Pow`
)

type Response struct {
	Challenge string `json:"challenge"`
}
