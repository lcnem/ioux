package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/iou module sentinel errors
var (
	ErrInvalidNamespaceId = sdkerrors.Register(ModuleName, 1, "Invalid Namespace ID")
	ErrInvalidPrefix      = sdkerrors.Register(ModuleName, 2, "Invalid Prefix")
)
