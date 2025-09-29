package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/storage module sentinel errors
var (
	ErrInvalidSigner        = errorsmod.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidOriginalHash  = errorsmod.Register(ModuleName, 1101, "invalid original hash")
	ErrFileAlreadyExists    = errorsmod.Register(ModuleName, 1102, "file with this hash already exists")
	ErrInvalidPacketTimeout = errorsmod.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = errorsmod.Register(ModuleName, 1501, "invalid version")
	ErrInvalidChannelFlow   = errorsmod.Register(ModuleName, 1502, "invalid channel flow")
)
