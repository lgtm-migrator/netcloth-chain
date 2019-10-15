package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/NetCloth/netcloth-chain/codec"
	"github.com/NetCloth/netcloth-chain/modules/ipal/types"
	sdk "github.com/NetCloth/netcloth-chain/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.

	// codespace
	codespace sdk.CodespaceType
}

// NewKeeper creates new instances of the nch Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		storeKey:  storeKey,
		cdc:       cdc,
		codespace: codespace,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// return the codespace
func (k Keeper) Codespace() sdk.CodespaceType {
	return k.codespace
}

// get a single ipal object
func (k Keeper) GetIPALObject(ctx sdk.Context, userAddress, serverIP string) (obj types.IPALObject, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetIPALObjectKey(userAddress))
	ctx.Logger().Info(string(types.GetIPALObjectKey(userAddress)))
	if value == nil {
		return obj, false
	}

	obj = types.MustUnmarshalIPALObject(k.cdc, value)
	return obj, true
}

// set ipal object
func (k Keeper) SetIPALObject(ctx sdk.Context, obj types.IPALObject) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalIPALObject(k.cdc, obj)
	store.Set(types.GetIPALObjectKey(obj.UserAddress), bz)
	ctx.Logger().Info(string(types.GetIPALObjectKey(obj.UserAddress)))
}