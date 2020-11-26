package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/foo/foo/x/foo/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) CreatePost(ctx sdk.Context, post types.Post) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(types.KeyPrefix(types.PostKey + post.Id), b)
}

func (k Keeper) UpdatePost(ctx sdk.Context, post types.Post) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := k.cdc.MustMarshalBinaryBare(&post)
	store.Set(types.KeyPrefix(types.PostKey + post.Id), b)
}

func (k Keeper) GetPost(ctx sdk.Context, key string) types.Post {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	var post types.Post
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostKey + key)), &post)
	return post
}

func (k Keeper) HasPost(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	return store.Has(types.KeyPrefix(types.PostKey + id))
}

func (k Keeper) GetPostOwner(ctx sdk.Context, key string) sdk.AccAddress {
    return k.GetPost(ctx, key).Creator
}

// DeletePost deletes a post
func (k Keeper) DeletePost(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	store.Delete(types.KeyPrefix(types.PostKey + key))
}

func (k Keeper) GetAllPost(ctx sdk.Context) (msgs []types.Post) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PostKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Post
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
