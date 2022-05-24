package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deepstack/auction/x/auction/types"
)

func (k msgServer) CreateObject(goCtx context.Context, msg *types.MsgCreateObject) (*types.MsgCreateObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var object = types.Object{
		Creator:          msg.Creator,
		Description:      msg.Description,
		Owner:            msg.Owner,
		Value:            msg.Value,
		PrevOwnSign:      msg.PrevOwnSign,
		StartBlockHeight: msg.StartBlockHeight,
		EndBlockHeight:   msg.EndBlockHeight,
		AuctionDuration:  msg.AuctionDuration,
		MarginBlocks:     msg.MarginBlocks,
		HighestBid:       msg.HighestBid,
		HighestBidder:    msg.HighestBidder,
	}

	id := k.AppendObject(
		ctx,
		object,
	)

	return &types.MsgCreateObjectResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateObject(goCtx context.Context, msg *types.MsgUpdateObject) (*types.MsgUpdateObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var object = types.Object{
		Creator:          msg.Creator,
		Id:               msg.Id,
		Description:      msg.Description,
		Owner:            msg.Owner,
		Value:            msg.Value,
		PrevOwnSign:      msg.PrevOwnSign,
		StartBlockHeight: msg.StartBlockHeight,
		EndBlockHeight:   msg.EndBlockHeight,
		AuctionDuration:  msg.AuctionDuration,
		MarginBlocks:     msg.MarginBlocks,
		HighestBid:       msg.HighestBid,
		HighestBidder:    msg.HighestBidder,
	}

	// Checks that the element exists
	val, found := k.GetObject(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetObject(ctx, object)

	return &types.MsgUpdateObjectResponse{}, nil
}

func (k msgServer) DeleteObject(goCtx context.Context, msg *types.MsgDeleteObject) (*types.MsgDeleteObjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetObject(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveObject(ctx, msg.Id)

	return &types.MsgDeleteObjectResponse{}, nil
}
