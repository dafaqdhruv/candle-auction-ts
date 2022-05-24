package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/deepstack/auction/testutil/keeper"
	"github.com/deepstack/auction/testutil/nullify"
	"github.com/deepstack/auction/x/auction/types"
)

func TestObjectQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNObject(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetObjectRequest
		response *types.QueryGetObjectResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetObjectRequest{Id: msgs[0].Id},
			response: &types.QueryGetObjectResponse{Object: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetObjectRequest{Id: msgs[1].Id},
			response: &types.QueryGetObjectResponse{Object: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetObjectRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Object(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestObjectQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNObject(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllObjectRequest {
		return &types.QueryAllObjectRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ObjectAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Object), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Object),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ObjectAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Object), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Object),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ObjectAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Object),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ObjectAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
