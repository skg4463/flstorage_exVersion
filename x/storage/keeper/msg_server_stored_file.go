package keeper

import (
	"context"
	"errors"
	"fmt"
	"crypto/sha256"
	"encoding/hex"

	"flstorage/x/storage/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStoredFile(ctx context.Context, msg *types.MsgCreateStoredFile) (*types.MsgCreateStoredFileResponse, error) {
	// --- 1.커스텀 검증 로직 ---
	_, err := hex.DecodeString(msg.OriginalHash)
	if err != nil || len(msg.OriginalHash) != sha256.Size*2 {
		// types/errors.go 에 정의한 커스텀 에러를 사용
		return nil, errorsmod.Wrapf(types.ErrInvalidOriginalHash, "해시: %s", msg.OriginalHash)
	}

	// --- 2. Ignite가 생성한 'collections'를 사용하여 중복 확인 ---
	// k.HasStoredFile(ctx, ...) 대신 k.StoredFile.Has(ctx, ...)를 사용
	ok, err := k.StoredFile.Has(ctx, msg.OriginalHash)
	if err != nil {
		// collections가 반환하는 에러를 그대로 래핑
		return nil, err
	}
	if ok {
		// 커스텀 에러
		return nil, errorsmod.Wrapf(types.ErrFileAlreadyExists, "해시: %s", msg.OriginalHash)
	}

	var storedFile = types.StoredFile{
		Creator:      msg.Creator,
		OriginalHash: msg.OriginalHash,
		Tag:          msg.Tag,
		ShardHashes:  msg.ShardHashes,
	}

	// --- 3. 'collections'를 사용하여 데이터 저장 ---
	// k.SetStoredFile(ctx, ...) 대신 k.StoredFile.Set(ctx, ...)를 사용합니다.
	if err := k.StoredFile.Set(ctx, storedFile.OriginalHash, storedFile); err != nil {
		return nil, err
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
