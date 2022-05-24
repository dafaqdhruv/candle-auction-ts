package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/deepstack/auction/x/auction/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-object [description] [owner] [value] [prev-own-sign] [start-block-height] [end-block-height] [auction-duration] [margin-blocks] [highest-bid] [highest-bidder]",
		Short: "Create a new object",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDescription := args[0]
			argOwner := args[1]
			argValue := args[2]
			argPrevOwnSign := args[3]
			argStartBlockHeight, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argEndBlockHeight, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argAuctionDuration, err := cast.ToUint64E(args[6])
			if err != nil {
				return err
			}
			argMarginBlocks, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}
			argHighestBid := args[8]
			argHighestBidder := args[9]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateObject(clientCtx.GetFromAddress().String(), argDescription, argOwner, argValue, argPrevOwnSign, argStartBlockHeight, argEndBlockHeight, argAuctionDuration, argMarginBlocks, argHighestBid, argHighestBidder)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-object [id] [description] [owner] [value] [prev-own-sign] [start-block-height] [end-block-height] [auction-duration] [margin-blocks] [highest-bid] [highest-bidder]",
		Short: "Update a object",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argDescription := args[1]

			argOwner := args[2]

			argValue := args[3]

			argPrevOwnSign := args[4]

			argStartBlockHeight, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			argEndBlockHeight, err := cast.ToUint64E(args[6])
			if err != nil {
				return err
			}

			argAuctionDuration, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}

			argMarginBlocks, err := cast.ToUint64E(args[8])
			if err != nil {
				return err
			}

			argHighestBid := args[9]

			argHighestBidder := args[10]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateObject(clientCtx.GetFromAddress().String(), id, argDescription, argOwner, argValue, argPrevOwnSign, argStartBlockHeight, argEndBlockHeight, argAuctionDuration, argMarginBlocks, argHighestBid, argHighestBidder)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteObject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-object [id]",
		Short: "Delete a object by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteObject(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
