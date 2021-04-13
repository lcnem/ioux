package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/ioux/types"
)

var _ sdk.Msg = &MsgCreateIou{}

func NewMsgCreateIou(namespace string, baseDenom string, issuer types.StringAccAddress, admin types.StringAccAddress) *MsgCreateIou {
	return &MsgCreateIou{
		Namespace: namespace, BaseDenom: baseDenom, Issuer: issuer, Admin: admin,
	}
}

func (msg *MsgCreateIou) Route() string {
	return RouterKey
}

func (msg *MsgCreateIou) Type() string {
	return "CreateIou"
}

func (msg *MsgCreateIou) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(string(msg.Admin))
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgCreateIou) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIou) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(string(msg.Issuer))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuer address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(string(msg.Admin))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateIouIssuer{}

func NewMsgUpdateIouIssuer(namespace string, baseDenom string, issuer types.StringAccAddress, admin types.StringAccAddress) *MsgUpdateIouIssuer {
	return &MsgUpdateIouIssuer{
		Namespace: namespace, BaseDenom: baseDenom, Issuer: issuer, Admin: admin,
	}
}

func (msg *MsgUpdateIouIssuer) Route() string {
	return RouterKey
}

func (msg *MsgUpdateIouIssuer) Type() string {
	return "CreateIou"
}

func (msg *MsgUpdateIouIssuer) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(string(msg.Admin))
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgUpdateIouIssuer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateIouIssuer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(string(msg.Issuer))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuer address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(string(msg.Admin))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgIssueIou{}

func NewMsgIssueIou(namespace string, prefix string, baseDenom string, issuer types.StringAccAddress, amount sdk.Coin) *MsgIssueIou {
	return &MsgIssueIou{
		Namespace: namespace, Prefix: prefix, BaseDenom: baseDenom, Issuer: issuer, Amount: amount,
	}
}

func (msg *MsgIssueIou) Route() string {
	return RouterKey
}

func (msg *MsgIssueIou) Type() string {
	return "UpdateIou"
}

func (msg *MsgIssueIou) GetSigners() []sdk.AccAddress {
	issuer, err := sdk.AccAddressFromBech32(string(msg.Issuer))
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{issuer}
}

func (msg *MsgIssueIou) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssueIou) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(string(msg.Issuer))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuer address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgBurnIou{}

func NewMsgBurnIou(namespace string, prefix string, baseDenom string, holder types.StringAccAddress, amount sdk.Coin) *MsgBurnIou {
	return &MsgBurnIou{
		Namespace: namespace, Prefix: prefix, BaseDenom: baseDenom, Holder: holder, Amount: amount,
	}
}
func (msg *MsgBurnIou) Route() string {
	return RouterKey
}

func (msg *MsgBurnIou) Type() string {
	return "DeleteIou"
}

func (msg *MsgBurnIou) GetSigners() []sdk.AccAddress {
	holder, err := sdk.AccAddressFromBech32(string(msg.Holder))
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{holder}
}

func (msg *MsgBurnIou) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnIou) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(string(msg.Holder))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid holder address (%s)", err)
	}
	return nil
}
