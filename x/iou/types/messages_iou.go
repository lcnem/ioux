package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/ioux/types"
)

var _ sdk.Msg = &MsgCreateIouNamespace{}

func NewMsgCreateIouNamespace(id string, admin sdk.AccAddress, issuer sdk.AccAddress) *MsgCreateIouNamespace {
	return &MsgCreateIouNamespace{
		Id: id, Admin: types.StringAccAddress(admin), Issuer: types.StringAccAddress(issuer),
	}
}

func (msg *MsgCreateIouNamespace) Route() string {
	return RouterKey
}

func (msg *MsgCreateIouNamespace) Type() string {
	return "CreateIou"
}

func (msg *MsgCreateIouNamespace) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(string(msg.Admin))
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgCreateIouNamespace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIouNamespace) ValidateBasic() error {
	if !namespaceRegExp.MatchString(msg.Id) {
		return ErrInvalidNamespaceId
	}
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

var _ sdk.Msg = &MsgUpdateIouNamespaceAdmin{}

func NewMsgUpdateIouNamespaceAdmin(id string, adminBefore sdk.AccAddress, adminAfter sdk.AccAddress) *MsgUpdateIouNamespaceAdmin {
	return &MsgUpdateIouNamespaceAdmin{
		Id: id, AdminBefore: types.StringAccAddress(adminBefore), AdminAfter: types.StringAccAddress(adminAfter),
	}
}

func (msg *MsgUpdateIouNamespaceAdmin) Route() string {
	return RouterKey
}

func (msg *MsgUpdateIouNamespaceAdmin) Type() string {
	return "CreateIou"
}

func (msg *MsgUpdateIouNamespaceAdmin) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(string(msg.AdminBefore))
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgUpdateIouNamespaceAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateIouNamespaceAdmin) ValidateBasic() error {
	if !namespaceRegExp.MatchString(msg.Id) {
		return ErrInvalidNamespaceId
	}
	_, err := sdk.AccAddressFromBech32(string(msg.AdminBefore))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(string(msg.AdminAfter))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateIouNamespaceIssuer{}

func NewMsgUpdateIouNamespaceIssuer(id string, admin sdk.AccAddress, issuer sdk.AccAddress) *MsgUpdateIouNamespaceIssuer {
	return &MsgUpdateIouNamespaceIssuer{
		Id: id, Admin: types.StringAccAddress(admin), Issuer: types.StringAccAddress(issuer),
	}
}

func (msg *MsgUpdateIouNamespaceIssuer) Route() string {
	return RouterKey
}

func (msg *MsgUpdateIouNamespaceIssuer) Type() string {
	return "CreateIou"
}

func (msg *MsgUpdateIouNamespaceIssuer) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(string(msg.Admin))
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

func (msg *MsgUpdateIouNamespaceIssuer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateIouNamespaceIssuer) ValidateBasic() error {
	if !namespaceRegExp.MatchString(msg.Id) {
		return ErrInvalidNamespaceId
	}
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

func NewMsgIssueIou(namespaceId string, prefix string, baseDenom string, issuer sdk.AccAddress, amount sdk.Int) *MsgIssueIou {
	return &MsgIssueIou{
		NamespaceId: namespaceId, Prefix: prefix, BaseDenom: baseDenom, Issuer: types.StringAccAddress(issuer), Amount: amount,
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
	if !namespaceRegExp.MatchString(msg.NamespaceId) {
		return ErrInvalidNamespaceId
	}
	if !prefixRegExp.MatchString(msg.Prefix) {
		return ErrInvalidPrefix
	}
	denom := GetDenom(msg.NamespaceId, msg.Prefix, msg.BaseDenom)
	if sdk.ValidateDenom(denom) != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid denom %s", denom)
	}

	_, err := sdk.AccAddressFromBech32(string(msg.Issuer))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuer address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(string(msg.Destination))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuer address (%s)", err)
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount must be positive")
	}
	return nil
}

var _ sdk.Msg = &MsgBurnIou{}

func NewMsgBurnIou(namespaceId string, prefix string, baseDenom string, holder sdk.AccAddress, amount sdk.Int) *MsgBurnIou {
	return &MsgBurnIou{
		NamespaceId: namespaceId, Prefix: prefix, BaseDenom: baseDenom, Holder: types.StringAccAddress(holder), Amount: amount,
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
	if !namespaceRegExp.MatchString(msg.NamespaceId) {
		return ErrInvalidNamespaceId
	}
	if !prefixRegExp.MatchString(msg.Prefix) {
		return ErrInvalidPrefix
	}
	denom := GetDenom(msg.NamespaceId, msg.Prefix, msg.BaseDenom)
	if sdk.ValidateDenom(denom) != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid denom %s", denom)
	}

	_, err := sdk.AccAddressFromBech32(string(msg.Holder))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid holder address (%s)", err)
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "amount must be positive")
	}

	return nil
}
