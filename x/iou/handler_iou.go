package iou

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/ioux/x/iou/keeper"
	"github.com/lcnem/ioux/x/iou/types"
)

func handleMsgCreateIouNamespace(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateIouNamespace) (*sdk.Result, error) {
	// Checks that the element exists
	if k.HasIouNamespace(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	k.CreateIouNamespace(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateIouNamespaceAdmin(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateIouNamespaceAdmin) (*sdk.Result, error) {

	// Checks that the element exists
	if !k.HasIouNamespace(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if string(msg.AdminBefore) != string(k.GetIouNamespaceAdmin(ctx, msg.Id)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var iou = types.IouNamespace{
		Id:     msg.Id,
		Admin:  msg.AdminAfter,
		Issuer: k.GetIouNamespaceIssuer(ctx, msg.Id),
	}

	k.SetIouNamespace(ctx, iou)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateIouNamespaceIssuer(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateIouNamespaceIssuer) (*sdk.Result, error) {
	var iou = types.IouNamespace{
		Id:     msg.Id,
		Admin:  msg.Admin,
		Issuer: msg.Issuer,
	}

	// Checks that the element exists
	if !k.HasIouNamespace(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if string(msg.Admin) != string(k.GetIouNamespaceAdmin(ctx, msg.Id)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect admin")
	}

	k.SetIouNamespace(ctx, iou)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgIssueIou(ctx sdk.Context, k keeper.Keeper, msg *types.MsgIssueIou) (*sdk.Result, error) {

	// Checks that the element exists
	if !k.HasIouNamespace(ctx, msg.NamespaceId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.NamespaceId))
	}

	// Checks if the the msg sender is the same as the current owner
	if string(msg.Issuer) != string(k.GetIouNamespaceIssuer(ctx, msg.NamespaceId)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	denom := types.GetDenom(msg.NamespaceId, msg.Prefix, msg.BaseDenom)
	coin := sdk.NewCoin(denom, msg.Amount)
	k.IssueIou(ctx, coin, msg.Issuer.AccAddress())

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgBurnIou(ctx sdk.Context, k keeper.Keeper, msg *types.MsgBurnIou) (*sdk.Result, error) {
	if !k.HasIouNamespace(ctx, msg.NamespaceId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.NamespaceId))
	}

	denom := types.GetDenom(msg.NamespaceId, msg.Prefix, msg.BaseDenom)
	coin := sdk.NewCoin(denom, msg.Amount)
	k.BurnIou(ctx, coin, msg.Holder.AccAddress())

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
