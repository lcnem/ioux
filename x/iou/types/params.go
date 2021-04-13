package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
)

var emptyDec = sdk.Dec{}

// Defaults for auction params
const ()

var ()

var _ paramstype.ParamSet = &Params{}

// NewParams returns a new Params object.
func NewParams() Params {
	return Params{}
}

// DefaultParams returns the default parameters for auctions.
func DefaultParams() Params {
	return NewParams()
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs.
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{}
}

// Equal returns a boolean determining if two Params types are identical.
func (p Params) Equal(p2 Params) bool {
	bz1 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p)
	bz2 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p2)
	return bytes.Equal(bz1, bz2)
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	return nil
}
