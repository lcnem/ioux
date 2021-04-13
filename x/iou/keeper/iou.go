package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/ioux/x/iou/types"
)

// GetIouCount get the total number of iou
func (k Keeper) GetIouCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouCountKey))
	byteKey := types.KeyPrefix(types.IouCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetIouCount set the total number of iou
func (k Keeper) SetIouCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouCountKey))
	byteKey := types.KeyPrefix(types.IouCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateIou creates a iou with a new id and update the count
func (k Keeper) CreateIou(ctx sdk.Context, msg types.MsgCreateIou) {
	// Create the iou
	count := k.GetIouCount(ctx)
	var iou = types.Iou{
		Creator: msg.Creator,
		Id:      strconv.FormatInt(count, 10),
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouKey))
	key := types.KeyPrefix(types.IouKey + iou.Id)
	value := k.cdc.MustMarshalBinaryBare(&iou)
	store.Set(key, value)

	// Update iou count
	k.SetIouCount(ctx, count+1)
}

// SetIou set a specific iou in the store
func (k Keeper) SetIou(ctx sdk.Context, iou types.Iou) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouKey))
	b := k.cdc.MustMarshalBinaryBare(&iou)
	store.Set(types.KeyPrefix(types.IouKey+iou.Id), b)
}

// GetIou returns a iou from its id
func (k Keeper) GetIou(ctx sdk.Context, key string) types.Iou {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouKey))
	var iou types.Iou
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.IouKey+key)), &iou)
	return iou
}

// HasIou checks if the iou exists
func (k Keeper) HasIou(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouKey))
	return store.Has(types.KeyPrefix(types.IouKey + id))
}

// GetIouOwner returns the creator of the iou
func (k Keeper) GetIouOwner(ctx sdk.Context, key string) string {
	return k.GetIou(ctx, key).Creator
}

// DeleteIou deletes a iou
func (k Keeper) DeleteIou(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouKey))
	store.Delete(types.KeyPrefix(types.IouKey + key))
}

// GetAllIou returns all iou
func (k Keeper) GetAllIou(ctx sdk.Context) (msgs []types.Iou) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.IouKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Iou
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
