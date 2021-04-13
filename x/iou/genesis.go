package iou

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/ioux/x/iou/keeper"
	"github.com/lcnem/ioux/x/iou/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the iou
	for _, elem := range genState.IouList {
		k.SetIou(ctx, *elem)
	}

	// Set iou count
	k.SetIouCount(ctx, int64(len(genState.IouList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all iou
	iouList := k.GetAllIou(ctx)
	for _, elem := range iouList {
		elem := elem
		genesis.IouList = append(genesis.IouList, &elem)
	}

	return genesis
}
