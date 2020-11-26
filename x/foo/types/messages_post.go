package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePost{}

func NewMsgCreatePost(creator sdk.AccAddress, title string, body string) *MsgCreatePost {
  return &MsgCreatePost{
		Creator: creator,
    Title: title,
    Body: body,
	}
}

func (msg *MsgCreatePost) Route() string {
  return RouterKey
}

func (msg *MsgCreatePost) Type() string {
  return "CreatePost"
}

func (msg *MsgCreatePost) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgCreatePost) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePost) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}

var _ sdk.Msg = &MsgUpdatePost{}

func NewMsgUpdatePost(creator sdk.AccAddress, id string, title string, body string) *MsgUpdatePost {
  return &MsgUpdatePost{
        Id: id,
		Creator: creator,
    Title: title,
    Body: body,
	}
}

func (msg *MsgUpdatePost) Route() string {
  return RouterKey
}

func (msg *MsgUpdatePost) Type() string {
  return "UpdatePost"
}

func (msg *MsgUpdatePost) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgUpdatePost) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePost) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}

var _ sdk.Msg = &MsgCreatePost{}

func NewMsgDeletePost(creator sdk.AccAddress, id string) *MsgDeletePost {
  return &MsgDeletePost{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeletePost) Route() string {
  return RouterKey
}

func (msg *MsgDeletePost) Type() string {
  return "DeletePost"
}

func (msg *MsgDeletePost) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgDeletePost) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePost) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
