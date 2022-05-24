package keeper

import (
	"github.com/deepstack/auction/x/auction/types"
)

var _ types.QueryServer = Keeper{}
