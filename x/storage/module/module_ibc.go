package storage

import (
    errorsmod "cosmossdk.io/errors"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
    ibcexported "github.com/cosmos/ibc-go/v10/modules/core/exported"

    "flstorage/x/storage/keeper"
    "flstorage/x/storage/types"
)

// IBCModule implements the ICS26 interface for storage module
type IBCModule struct {
    cdc    codec.Codec
    keeper keeper.Keeper
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(cdc codec.Codec, keeper keeper.Keeper) IBCModule {
    return IBCModule{
        cdc:    cdc,
        keeper: keeper,
    }
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
    ctx sdk.Context,
    order channeltypes.Order,
    connectionHops []string,
    portID string,
    channelID string,
    counterparty channeltypes.Counterparty,
    version string,
) (string, error) {
    return types.Version, nil
}

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
    ctx sdk.Context,
    order channeltypes.Order,
    connectionHops []string,
    portID,
    channelID string,
    counterparty channeltypes.Counterparty,
    counterpartyVersion string,
) (version string, err error) {
    return types.Version, nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
    ctx sdk.Context,
    portID,
    channelID string,
    counterpartyChannelID string,
    counterpartyVersion string,
) error {
    return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
    ctx sdk.Context,
    portID,
    channelID string,
) error {
    return nil
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
    ctx sdk.Context,
    portID,
    channelID string,
) error {
    return errorsmod.Wrap(types.ErrInvalidChannelFlow, "cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
    ctx sdk.Context,
    portID,
    channelID string,
) error {
    return nil
}

// OnRecvPacket implements the IBCModule interface
func (im IBCModule) OnRecvPacket(
    ctx sdk.Context,
    channelID string,        // 추가된 파라미터
    packet channeltypes.Packet,
    relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
    return channeltypes.NewResultAcknowledgement([]byte{byte(1)})
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
    ctx sdk.Context,
    portID string,
    packet channeltypes.Packet,
    acknowledgement []byte,
    relayer sdk.AccAddress,
) error {
    return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCModule) OnTimeoutPacket(
    ctx sdk.Context,
    portID string,
    packet channeltypes.Packet,
    relayer sdk.AccAddress,
) error {
    return nil
}