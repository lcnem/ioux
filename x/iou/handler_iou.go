package iou

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/ioux/x/iou/keeper"
	"github.com/lcnem/ioux/x/iou/types"
)

func handleMsgCreateIou(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateIou) (*sdk.Result, error) {
	k.CreateIou(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateIou(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateIou) (*sdk.Result, error) {
	var iou = types.Iou{
		Creator: msg.Creator,
		Id:      msg.Id,
	}

	// Checks that the element exists
	if !k.HasIou(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetIouOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetIou(ctx, iou)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteIou(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteIou) (*sdk.Result, error) {
	if !k.HasIou(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetIouOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.DeleteIou(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
