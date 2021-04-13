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

func (k Keeper) IouAll(c context.Context, req *types.QueryAllIouRequest) (*types.QueryAllIouResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ious []*types.Iou
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	iouStore := prefix.NewStore(store, types.KeyPrefix(types.IouKey))

	pageRes, err := query.Paginate(iouStore, req.Pagination, func(key []byte, value []byte) error {
		var iou types.Iou
		if err := k.cdc.UnmarshalBinaryBare(value, &iou); err != nil {
			return err
		}

		ious = append(ious, &iou)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIouResponse{Iou: ious, Pagination: pageRes}, nil
}

func (k Keeper) Iou(c context.Context, req *types.QueryGetIouRequest) (*types.QueryGetIouResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var iou types.Iou
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.IouKey+req.Id)), &iou)

	return &types.QueryGetIouResponse{Iou: &iou}, nil
}
