package OpenseaAPI

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zeebo/errs"
)

type RetrieveEventsRequest struct {
	OnlyOpenSea          bool   `json:"only_opensea"`
	TokenID              int32  `json:"token_id"`
	AssetContractAddress string `json:"asset_contract_address"`
	CollectionSlug       string `json:"collection_slug"`
	AccountAddress       string `json:"account_address"`
	EventType            string `json:"event_type"`
	AuctionType          string `json:"auction_type"`
	OccurredBefore       string `json:"occurred_before"`
	OccurredAfter        string `json:"occurred_after"`
	Cursor               string `json:"cursor"`
}

type RetrieveEventsResponse struct {
	Next        string       `json:"next"`
	Previous    string       `json:"previous"`
	AssetEvents []AssetEvent `json:"asset_events"`
}

func (client *Client) RetrieveEvents(ctx context.Context, r RetrieveEventsRequest) (RetrieveEventsResponse, error) {
	url, err := makeURL(r)
	if err != nil {
		return RetrieveEventsResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+url, nil)
	if err != nil {
		return RetrieveEventsResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "c9128ae930224cabac7af252a18759a1")

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return RetrieveEventsResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		rr, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return RetrieveEventsResponse{}, err
		}

		return RetrieveEventsResponse{}, errs.New(string(rr))
	}

	var retrieveEventsResponse RetrieveEventsResponse
	if err = json.NewDecoder(resp.Body).Decode(&retrieveEventsResponse); err != nil {
		return RetrieveEventsResponse{}, err
	}

	return retrieveEventsResponse, nil
}

func makeURL(r RetrieveEventsRequest) (url string, _ error) {
	url = "events"

	if strconv.Itoa(int(r.TokenID)) != "0" && r.AssetContractAddress == "" {
		return "", fmt.Errorf("asset contract address required if token ID passed")
	}

	if r.OnlyOpenSea {
		url = url + "?only_opensea=true"
	} else {
		url = url + "?only_opensea=false"
	}

	if r.AuctionType != "" {
		switch r.AuctionType {
		case "english":
			url = url + "&auction_type=" + r.AuctionType
		case "dutch":
			url = url + "&auction_type=" + r.AuctionType
		case "min_price":
			url = url + "&auction_type=" + r.AuctionType
		default:
			return "", fmt.Errorf("wrong format of Auction Type")
		}
	}

	if r.EventType != "" {
		switch r.EventType {
		case "created":
			url = url + "&event_type=" + r.EventType
		case "successful":
			url = url + "&event_type=" + r.EventType
		case "cancelled":
			url = url + "&event_type=" + r.EventType
		case "bid_entered":
			url = url + "&event_type=" + r.EventType
		case "bid_withdrawn":
			url = url + "&event_type=" + r.EventType
		case "transfer":
			url = url + "&event_type=" + r.EventType
		case "approve":
			url = url + "&event_type=" + r.EventType
		default:
			return "", fmt.Errorf("wrong format of Event Type")
		}
	}

	if r.CollectionSlug != "" {
		url = url + "&collection_slug=" + r.CollectionSlug
	}

	if r.TokenID != 0 {
		tokenID := strconv.Itoa(int(r.TokenID))
		url = url + "&token_id=" + tokenID
	}

	if r.AssetContractAddress != "" {
		url = url + "&asset_contract_address=" + r.AssetContractAddress
	}

	if r.Cursor != "" {
		url = url + "&cursor=" + r.Cursor
	}

	if r.AccountAddress != "" {
		url = url + "&account_address=" + r.AccountAddress
	}

	if r.OccurredAfter != "" {
		url = url + "&occurred_after=" + r.OccurredAfter
	}

	if r.OccurredBefore != "" {
		url = url + "&occurred_before=" + r.OccurredBefore
	}

	return url, nil
}
