// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package opensea

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zeebo/errs"
)

type SingleContractRequest struct {
	AssetContractAddress string `json:"asset_contract_address"`
}

type SingleContractResponse struct {
	Collection                  Collection  `json:"collection"`
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       int         `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageUrl                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}

func (client *Client) SingleContract(ctx context.Context, r SingleContractRequest) (SingleContractResponse, error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"asset_contract/"+r.AssetContractAddress, nil)
	if err != nil {
		return SingleContractResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", client.config.APIKey)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return SingleContractResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		rr, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return SingleContractResponse{}, err
		}

		return SingleContractResponse{}, errs.New(string(rr))
	}

	var singleContractResponse SingleContractResponse
	if err = json.NewDecoder(resp.Body).Decode(&singleContractResponse); err != nil {
		return SingleContractResponse{}, err
	}

	return singleContractResponse, nil

}
