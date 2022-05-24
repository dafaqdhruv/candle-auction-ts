package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/deepstack/auction/testutil/keeper"
	"github.com/deepstack/auction/testutil/nullify"
	"github.com/deepstack/auction/x/auction/keeper"
	"github.com/deepstack/auction/x/auction/types"
	"github.com/stretchr/testify/require"
)

func createNObject(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Object {
	items := make([]types.Object, n)
	for i := range items {
		items[i].Id = keeper.AppendObject(ctx, items[i])
	}
	return items
}

func TestObjectGet(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNObject(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetObject(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestObjectRemove(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNObject(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveObject(ctx, item.Id)
		_, found := keeper.GetObject(ctx, item.Id)
		require.False(t, found)
	}
}

func TestObjectGetAll(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNObject(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllObject(ctx)),
	)
}

func TestObjectCount(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNObject(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetObjectCount(ctx))
}
