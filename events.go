// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package opensea

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zeebo/errs"
)

// ErrTooManyRequests indicates OpenSea API error when request was throttled.
var ErrTooManyRequests = errs.Class("Request was throttled")

type EventsRequest struct {
	OnlyOpenSea          bool   `json:"only_opensea"`
	TokenID              int64  `json:"token_id"`
	AssetContractAddress string `json:"asset_contract_address"`
	CollectionSlug       string `json:"collection_slug"`
	AccountAddress       string `json:"account_address"`
	EventType            string `json:"event_type"`
	AuctionType          string `json:"auction_type"`
	OccurredBefore       string `json:"occurred_before"`
	OccurredAfter        string `json:"occurred_after"`
	Cursor               string `json:"cursor"`
}

type EventsResponse struct {
	Next        string       `json:"next"`
	Previous    string       `json:"previous"`
	AssetEvents []AssetEvent `json:"asset_events"`
}

func (client *Client) Events(ctx context.Context, r EventsRequest) (EventsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"events", nil)
	if err != nil {
		return EventsResponse{}, err
	}

	req = setURL(req, r)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", client.config.APIKey)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return EventsResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		rr, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return EventsResponse{}, err
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			return EventsResponse{}, ErrTooManyRequests.Wrap(errs.New(string(rr)))
		}

		return EventsResponse{}, errs.New(string(rr))
	}

	var eventsResponse EventsResponse
	if err = json.NewDecoder(resp.Body).Decode(&eventsResponse); err != nil {
		return EventsResponse{}, err
	}

	return eventsResponse, nil
}

func setURL(req *http.Request, r EventsRequest) *http.Request {
	params := req.URL.Query()

	if r.OnlyOpenSea {
		params.Set("only_opensea", "true")
	} else {
		params.Set("only_opensea", "false")
	}

	if r.AuctionType != "" {
		params.Set("auction_type", r.AuctionType)
	}

	if r.EventType != "" {
		params.Set("event_type", r.EventType)
	}

	if r.CollectionSlug != "" {
		params.Set("collection_slug", r.CollectionSlug)
	}

	if r.TokenID != 0 {
		params.Set("token_id", strconv.Itoa(int(r.TokenID)))
	}

	if r.AssetContractAddress != "" {
		params.Set("asset_contract_address", r.AssetContractAddress)
	}

	if r.Cursor != "" {
		params.Set("cursor", r.Cursor)
	}

	if r.AccountAddress != "" {
		params.Set("account_address", r.AccountAddress)
	}

	if r.OccurredAfter != "" {
		params.Set("occurred_after", r.OccurredAfter)
	}

	if r.OccurredBefore != "" {
		params.Set("occurred_before", r.OccurredBefore)
	}

	req.URL.RawQuery = params.Encode()
	return req
}
