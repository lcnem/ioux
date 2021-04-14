package iou

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/ioux/x/iou/keeper"
	"github.com/lcnem/ioux/x/iou/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgCreateIouNamespace:
			return handleMsgCreateIouNamespace(ctx, k, msg)

		case *types.MsgUpdateIouNamespaceAdmin:
			return handleMsgUpdateIouNamespaceAdmin(ctx, k, msg)

		case *types.MsgUpdateIouNamespaceIssuer:
			return handleMsgUpdateIouNamespaceIssuer(ctx, k, msg)

		case *types.MsgIssueIou:
			return handleMsgIssueIou(ctx, k, msg)

		case *types.MsgBurnIou:
			return handleMsgBurnIou(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
