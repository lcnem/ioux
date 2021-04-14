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
	for _, elem := range genState.IouNamespaces {
		k.SetIouNamespace(ctx, *elem)
	}

	// Set iou count
	k.SetIouNamespaceCount(ctx, int64(len(genState.IouNamespaces)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all iou
	iouList := k.GetAllIouNamespace(ctx)
	for _, elem := range iouList {
		elem := elem
		genesis.IouNamespaces = append(genesis.IouNamespaces, &elem)
	}

	return genesis
}
