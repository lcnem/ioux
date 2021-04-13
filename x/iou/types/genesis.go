package types

import "fmt"

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: Params{},
		// this line is used by starport scaffolding # genesis/types/default
		Ious: []*Iou{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in iou
	iouIdMap := make(map[string]bool)

	for _, elem := range gs.Ious {
		if _, ok := iouIdMap[elem.Namespace+elem.BaseDenom]; ok {
			return fmt.Errorf("duplicated id for iou")
		}
		iouIdMap[elem.Namespace+elem.BaseDenom] = true
	}

	return nil
}
