package OpenseaAPI_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"OpenseaAPI"
)

func TestOpenseaAPI(t *testing.T) {
	client := OpenseaAPI.New(OpenseaAPI.Config{
		DefaultURL: "https://api.opensea.io/api/v1/",
		APIKey:     "c9128ae930224cabac7af252a18759a1",
	})

	resp, err := client.Collections(context.Background(), OpenseaAPI.CollectionStatsRequest{CollectionSlug: "doodles-official"})
	require.NoError(t, err)
	require.NotNil(t, resp)

	resp2, err := client.Events(context.Background(), OpenseaAPI.EventsRequest{
		OnlyOpenSea: true,
		EventType:   "successful",
	})
	require.NoError(t, err)
	require.NotNil(t, resp2)

	_, err = client.Events(context.Background(), OpenseaAPI.EventsRequest{
		OnlyOpenSea: true,
		TokenID:     123,
	})
	require.Error(t, err)

	_, err = client.Events(context.Background(), OpenseaAPI.EventsRequest{
		OnlyOpenSea: true,
		AuctionType: "random",
	})
	require.Error(t, err)

	_, err = client.Events(context.Background(), OpenseaAPI.EventsRequest{
		OnlyOpenSea: true,
		EventType:   "random",
	})
	require.Error(t, err)
}
