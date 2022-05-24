package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/deepstack/auction/x/auction/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ObjectAll(c context.Context, req *types.QueryAllObjectRequest) (*types.QueryAllObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var objects []types.Object
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	objectStore := prefix.NewStore(store, types.KeyPrefix(types.ObjectKey))

	pageRes, err := query.Paginate(objectStore, req.Pagination, func(key []byte, value []byte) error {
		var object types.Object
		if err := k.cdc.Unmarshal(value, &object); err != nil {
			return err
		}

		objects = append(objects, object)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllObjectResponse{Object: objects, Pagination: pageRes}, nil
}

func (k Keeper) Object(c context.Context, req *types.QueryGetObjectRequest) (*types.QueryGetObjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	object, found := k.GetObject(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetObjectResponse{Object: object}, nil
}
