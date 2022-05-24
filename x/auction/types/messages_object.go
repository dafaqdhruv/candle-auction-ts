package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateObject = "create_object"
	TypeMsgUpdateObject = "update_object"
	TypeMsgDeleteObject = "delete_object"
)

var _ sdk.Msg = &MsgCreateObject{}

func NewMsgCreateObject(creator string, description string, owner string, value string, prevOwnSign string, startBlockHeight uint64, endBlockHeight uint64, auctionDuration uint64, marginBlocks uint64, highestBid string, highestBidder string) *MsgCreateObject {
	return &MsgCreateObject{
		Creator:          creator,
		Description:      description,
		Owner:            owner,
		Value:            value,
		PrevOwnSign:      prevOwnSign,
		StartBlockHeight: startBlockHeight,
		EndBlockHeight:   endBlockHeight,
		AuctionDuration:  auctionDuration,
		MarginBlocks:     marginBlocks,
		HighestBid:       highestBid,
		HighestBidder:    highestBidder,
	}
}

func (msg *MsgCreateObject) Route() string {
	return RouterKey
}

func (msg *MsgCreateObject) Type() string {
	return TypeMsgCreateObject
}

func (msg *MsgCreateObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateObject{}

func NewMsgUpdateObject(creator string, id uint64, description string, owner string, value string, prevOwnSign string, startBlockHeight uint64, endBlockHeight uint64, auctionDuration uint64, marginBlocks uint64, highestBid string, highestBidder string) *MsgUpdateObject {
	return &MsgUpdateObject{
		Id:               id,
		Creator:          creator,
		Description:      description,
		Owner:            owner,
		Value:            value,
		PrevOwnSign:      prevOwnSign,
		StartBlockHeight: startBlockHeight,
		EndBlockHeight:   endBlockHeight,
		AuctionDuration:  auctionDuration,
		MarginBlocks:     marginBlocks,
		HighestBid:       highestBid,
		HighestBidder:    highestBidder,
	}
}

func (msg *MsgUpdateObject) Route() string {
	return RouterKey
}

func (msg *MsgUpdateObject) Type() string {
	return TypeMsgUpdateObject
}

func (msg *MsgUpdateObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteObject{}

func NewMsgDeleteObject(creator string, id uint64) *MsgDeleteObject {
	return &MsgDeleteObject{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteObject) Route() string {
	return RouterKey
}

func (msg *MsgDeleteObject) Type() string {
	return TypeMsgDeleteObject
}

func (msg *MsgDeleteObject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteObject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteObject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
