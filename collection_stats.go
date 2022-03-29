package opensea

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zeebo/errs"
)

type CollectionStatsRequest struct {
	CollectionSlug string `json:"collection_slug"`
}

type CollectionStatsResponse struct {
	Stats Stats `json:"stats"`
}

func (client *Client) CollectionStats(ctx context.Context, r CollectionStatsRequest) (CollectionStatsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"collection/"+r.CollectionSlug+"/stats", nil)
	if err != nil {
		return CollectionStatsResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", client.config.APIKey)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return CollectionStatsResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		rr, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return CollectionStatsResponse{}, err
		}

		return CollectionStatsResponse{}, errs.New(string(rr))
	}

	var collectionStatsResponse CollectionStatsResponse
	if err = json.NewDecoder(resp.Body).Decode(&collectionStatsResponse); err != nil {
		return CollectionStatsResponse{}, err
	}

	return collectionStatsResponse, nil
}
