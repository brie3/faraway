package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	gopow "github.com/bwesterb/go-pow"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/brie3/faraway/pkg/config"
	"github.com/brie3/faraway/pkg/model"
)

func main() {
	log.Logger = config.MustGet().Logger()

	ctx := context.Background()
	b := doRequest(ctx, nil)

	var response model.Response
	if err := json.Unmarshal(b, &response); err != nil {
		panic(err)
	}

	pow, err := gopow.Fulfil(response.Challenge, nil)
	if err != nil {
		panic(err)
	}

	b = doRequest(ctx, map[string]string{model.ChallengeKey: response.Challenge, model.POWKey: pow})
	log.Logger.Info().Bytes("body", b).Send()
}

func doRequest(ctx context.Context, header map[string]string) []byte {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://%s/quote", config.MustGet().ServiceBind), nil)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(errors.Errorf("status code: %d", resp.StatusCode))
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(errors.Wrap(err, "cannot read response body"))
	}
	return b
}
