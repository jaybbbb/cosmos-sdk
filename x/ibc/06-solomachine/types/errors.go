package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	SubModuleName = "solo machine"
)

var (
	ErrInvalidHeader   = sdkerrors.Register(SubModuleName, 2, "invalid header")
	ErrInvalidSequence = sdkerrors.Register(SubModuleName, 3, "invalid sequence")
)
