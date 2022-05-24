package keeper_test

import (
	"testing"

	testkeeper "github.com/deepstack/auction/testutil/keeper"
	"github.com/deepstack/auction/x/auction/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AuctionKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
