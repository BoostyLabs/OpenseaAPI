// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package opensea

import "context"

// Opensea interface to call OpenSea API.
type Opensea interface {
	CollectionStats(ctx context.Context, request CollectionStatsRequest) (CollectionStatsResponse, error)
	Events(ctx context.Context, request EventsRequest) (EventsResponse, error)
	SingleContract(ctx context.Context, request SingleContractRequest) (SingleContractResponse, error)
}
