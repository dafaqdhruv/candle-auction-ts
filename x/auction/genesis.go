package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deepstack/auction/x/auction/keeper"
	"github.com/deepstack/auction/x/auction/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the object
	for _, elem := range genState.ObjectList {
		k.SetObject(ctx, elem)
	}

	// Set object count
	k.SetObjectCount(ctx, genState.ObjectCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ObjectList = k.GetAllObject(ctx)
	genesis.ObjectCount = k.GetObjectCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
