package keeper

import (
	"context"
	"errors"
	"fmt"

	"flstorage/x/storage/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStoredFile(ctx context.Context, msg *types.MsgCreateStoredFile) (*types.MsgCreateStoredFileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.StoredFile.Has(ctx, msg.OriginalHash)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var storedFile = types.StoredFile{
		Creator:      msg.Creator,
		OriginalHash: msg.OriginalHash,
		Tag:          msg.Tag,
		ShardHashes:  msg.ShardHashes,
	}

	if err := k.StoredFile.Set(ctx, storedFile.OriginalHash, storedFile); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateStoredFileResponse{}, nil
}

func (k msgServer) UpdateStoredFile(ctx context.Context, msg *types.MsgUpdateStoredFile) (*types.MsgUpdateStoredFileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.StoredFile.Get(ctx, msg.OriginalHash)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var storedFile = types.StoredFile{
		Creator:      msg.Creator,
		OriginalHash: msg.OriginalHash,
		Tag:          msg.Tag,
		ShardHashes:  msg.ShardHashes,
	}

	if err := k.StoredFile.Set(ctx, storedFile.OriginalHash, storedFile); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update storedFile")
	}

	return &types.MsgUpdateStoredFileResponse{}, nil
}

func (k msgServer) DeleteStoredFile(ctx context.Context, msg *types.MsgDeleteStoredFile) (*types.MsgDeleteStoredFileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.StoredFile.Get(ctx, msg.OriginalHash)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.StoredFile.Remove(ctx, msg.OriginalHash); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove storedFile")
	}

	return &types.MsgDeleteStoredFileResponse{}, nil
}
