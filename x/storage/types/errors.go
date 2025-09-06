package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/storage module sentinel errors
var (
	ErrInvalidSigner = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	// --- 아래 두 줄 추가 ---
	ErrInvalidOriginalHash = errors.Register(ModuleName, 1101, "invalid original hash")
	ErrFileAlreadyExists   = errors.Register(ModuleName, 1102, "file with this hash already exists")
)
