package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/lcnem/ioux/x/iou/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IouNamespaceAll(c context.Context, req *types.QueryAllIouNamespaceRequest) (*types.QueryAllIouNamespaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ious []*types.IouNamespace
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	iouStore := prefix.NewStore(store, types.KeyPrefix(types.IouNamespaceKey))

	pageRes, err := query.Paginate(iouStore, req.Pagination, func(key []byte, value []byte) error {
		var iou types.IouNamespace
		if err := k.cdc.UnmarshalBinaryBare(value, &iou); err != nil {
			return err
		}

		ious = append(ious, &iou)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIouNamespaceResponse{IouNamespaces: ious, Pagination: pageRes}, nil
}

func (k Keeper) IouNamespace(c context.Context, req *types.QueryGetIouNamespaceRequest) (*types.QueryGetIouNamespaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var iou types.IouNamespace
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.IouNamespaceKey+req.Id)), &iou)

	return &types.QueryGetIouNamespaceResponse{IouNamepsace: &iou}, nil
}
