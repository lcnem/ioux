package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/lcnem/ioux/x/iou/types"
)

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		paramSpace    paramtypes.Subspace
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey, paramSpace paramtypes.Subspace, accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramSpace:    paramSpace,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) IssueIou(ctx sdk.Context, coin sdk.Coin, destination sdk.AccAddress) error {
	coins := sdk.NewCoins(coin)
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, destination, coins)
	if err != nil {
		k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
		return err
	}

	return nil
}

func (k Keeper) BurnIou(ctx sdk.Context, coin sdk.Coin, holder sdk.AccAddress) error {
	coins := sdk.NewCoins(coin)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, holder, types.ModuleName, coins)
	if err != nil {
		return err
	}
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)

	return nil
}
