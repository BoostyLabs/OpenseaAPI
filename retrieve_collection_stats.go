package OpenseaAPI

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zeebo/errs"
)

type RetrieveCollectionStatsRequest struct {
	CollectionSlug string `json:"collection_slug"`
}

type RetrieveCollectionStatsResponse struct {
	Stats Stats `json:"stats"`
}

func (client *Client) RetrieveCollections(ctx context.Context, r RetrieveCollectionStatsRequest) (RetrieveCollectionStatsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"collection/"+r.CollectionSlug+"/stats", nil)
	if err != nil {
		return RetrieveCollectionStatsResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "c9128ae930224cabac7af252a18759a1")

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return RetrieveCollectionStatsResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		rr, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return RetrieveCollectionStatsResponse{}, err
		}

		return RetrieveCollectionStatsResponse{}, errs.New(string(rr))
	}

	var retrieveCollectionStatsResponse RetrieveCollectionStatsResponse
	if err = json.NewDecoder(resp.Body).Decode(&retrieveCollectionStatsResponse); err != nil {
		return RetrieveCollectionStatsResponse{}, err
	}

	return retrieveCollectionStatsResponse, nil
}
