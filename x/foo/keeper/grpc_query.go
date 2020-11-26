package keeper

import (
	"github.com/foo/foo/x/foo/types"
)

var _ types.QueryServer = Keeper{}
