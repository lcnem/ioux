package keeper

import (
	"github.com/lcnem/ioux/x/iou/types"
)

var _ types.QueryServer = Keeper{}
