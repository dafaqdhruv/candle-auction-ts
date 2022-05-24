package auction_test

import (
	"testing"

	keepertest "github.com/deepstack/auction/testutil/keeper"
	"github.com/deepstack/auction/testutil/nullify"
	"github.com/deepstack/auction/x/auction"
	"github.com/deepstack/auction/x/auction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ObjectList: []types.Object{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ObjectCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AuctionKeeper(t)
	auction.InitGenesis(ctx, *k, genesisState)
	got := auction.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ObjectList, got.ObjectList)
	require.Equal(t, genesisState.ObjectCount, got.ObjectCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
