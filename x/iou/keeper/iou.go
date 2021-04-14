package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_types "github.com/lcnem/ioux/types"
	"github.com/lcnem/ioux/x/iou/types"
)

// GetIouNamespaceCount get the total number of IouNamespace
func (k Keeper) GetIouNamespaceCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceCountKey))
	byteKey := types.KeyPrefix(types.IouNamespaceCountKey)
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

// SetIouNamespaceCount set the total number of IouNamespace
func (k Keeper) SetIouNamespaceCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceCountKey))
	byteKey := types.KeyPrefix(types.IouNamespaceCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateIouNamespace creates a IouNamespace with a new id and update the count
func (k Keeper) CreateIouNamespace(ctx sdk.Context, msg types.MsgCreateIouNamespace) {
	// Create the IouNamespace
	count := k.GetIouNamespaceCount(ctx)
	var IouNamespace = types.IouNamespace{
		Id:     msg.Id,
		Admin:  msg.Admin,
		Issuer: msg.Issuer,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceKey))
	key := types.KeyPrefix(types.IouNamespaceKey + IouNamespace.Id)
	value := k.cdc.MustMarshalBinaryBare(&IouNamespace)
	store.Set(key, value)

	// Update IouNamespace count
	k.SetIouNamespaceCount(ctx, count+1)
}

// SetIouNamespace set a specific IouNamespace in the store
func (k Keeper) SetIouNamespace(ctx sdk.Context, IouNamespace types.IouNamespace) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceKey))
	b := k.cdc.MustMarshalBinaryBare(&IouNamespace)
	store.Set(types.KeyPrefix(types.IouNamespaceKey+IouNamespace.Id), b)
}

// GetIouNamespace returns a IouNamespace from its id
func (k Keeper) GetIouNamespace(ctx sdk.Context, id string) types.IouNamespace {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceKey))
	var IouNamespace types.IouNamespace
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.IouNamespaceKey+id)), &IouNamespace)
	return IouNamespace
}

// HasIouNamespace checks if the IouNamespace exists
func (k Keeper) HasIouNamespace(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceKey))
	return store.Has(types.KeyPrefix(types.IouNamespaceKey + id))
}

// GetIouNamespaceIssuer returns the issuer of the IouNamespace
func (k Keeper) GetIouNamespaceIssuer(ctx sdk.Context, id string) _types.StringAccAddress {
	return k.GetIouNamespace(ctx, id).Issuer
}

// GetIouNamespaceAdmin returns the admin of the IouNamespace
func (k Keeper) GetIouNamespaceAdmin(ctx sdk.Context, id string) _types.StringAccAddress {
	return k.GetIouNamespace(ctx, id).Admin
}

// DeleteIouNamespace deletes a IouNamespace
func (k Keeper) DeleteIouNamespace(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceKey))
	store.Delete(types.KeyPrefix(types.IouNamespaceKey + id))
}

// GetAllIouNamespace returns all IouNamespace
func (k Keeper) GetAllIouNamespace(ctx sdk.Context) (msgs []types.IouNamespace) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IouNamespaceKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.IouNamespaceKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.IouNamespace
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
