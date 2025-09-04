package storage

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"flstorage/testutil/sample"
	storagesimulation "flstorage/x/storage/simulation"
	"flstorage/x/storage/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	storageGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		StoredFileMap: []types.StoredFile{{Creator: sample.AccAddress(),
			OriginalHash: "0",
		}, {Creator: sample.AccAddress(),
			OriginalHash: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&storageGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateStoredFile          = "op_weight_msg_storage"
		defaultWeightMsgCreateStoredFile int = 100
	)

	var weightMsgCreateStoredFile int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateStoredFile, &weightMsgCreateStoredFile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStoredFile = defaultWeightMsgCreateStoredFile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStoredFile,
		storagesimulation.SimulateMsgCreateStoredFile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateStoredFile          = "op_weight_msg_storage"
		defaultWeightMsgUpdateStoredFile int = 100
	)

	var weightMsgUpdateStoredFile int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateStoredFile, &weightMsgUpdateStoredFile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStoredFile = defaultWeightMsgUpdateStoredFile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStoredFile,
		storagesimulation.SimulateMsgUpdateStoredFile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteStoredFile          = "op_weight_msg_storage"
		defaultWeightMsgDeleteStoredFile int = 100
	)

	var weightMsgDeleteStoredFile int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteStoredFile, &weightMsgDeleteStoredFile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStoredFile = defaultWeightMsgDeleteStoredFile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteStoredFile,
		storagesimulation.SimulateMsgDeleteStoredFile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
