package keeper

import (
	"github.com/test/planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}
