package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deepstack/auction/x/auction/types"
)

// GetObjectCount get the total number of object
func (k Keeper) GetObjectCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ObjectCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetObjectCount set the total number of object
func (k Keeper) SetObjectCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ObjectCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendObject appends a object in the store with a new id and update the count
func (k Keeper) AppendObject(
	ctx sdk.Context,
	object types.Object,
) uint64 {
	// Create the object
	count := k.GetObjectCount(ctx)

	// Set the ID of the appended value
	object.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObjectKey))
	appendedValue := k.cdc.MustMarshal(&object)
	store.Set(GetObjectIDBytes(object.Id), appendedValue)

	// Update object count
	k.SetObjectCount(ctx, count+1)

	return count
}

// SetObject set a specific object in the store
func (k Keeper) SetObject(ctx sdk.Context, object types.Object) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObjectKey))
	b := k.cdc.MustMarshal(&object)
	store.Set(GetObjectIDBytes(object.Id), b)
}

// GetObject returns a object from its id
func (k Keeper) GetObject(ctx sdk.Context, id uint64) (val types.Object, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObjectKey))
	b := store.Get(GetObjectIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveObject removes a object from the store
func (k Keeper) RemoveObject(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObjectKey))
	store.Delete(GetObjectIDBytes(id))
}

// GetAllObject returns all object
func (k Keeper) GetAllObject(ctx sdk.Context) (list []types.Object) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObjectKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Object
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetObjectIDBytes returns the byte representation of the ID
func GetObjectIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetObjectIDFromBytes returns ID in uint64 format from a byte array
func GetObjectIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
