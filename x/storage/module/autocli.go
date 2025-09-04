package storage

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"flstorage/x/storage/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListStoredFile",
					Use:       "list-stored-file",
					Short:     "List all StoredFile",
				},
				{
					RpcMethod:      "GetStoredFile",
					Use:            "get-stored-file [id]",
					Short:          "Gets a StoredFile",
					Alias:          []string{"show-stored-file"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "original_hash"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateStoredFile",
					Use:            "create-stored-file [original_hash] [tag] [shard-hashes]",
					Short:          "Create a new StoredFile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "original_hash"}, {ProtoField: "tag"}, {ProtoField: "shard_hashes", Varargs: true}},
				},
				{
					RpcMethod:      "UpdateStoredFile",
					Use:            "update-stored-file [original_hash] [tag] [shard-hashes]",
					Short:          "Update StoredFile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "original_hash"}, {ProtoField: "tag"}, {ProtoField: "shard_hashes", Varargs: true}},
				},
				{
					RpcMethod:      "DeleteStoredFile",
					Use:            "delete-stored-file [original_hash]",
					Short:          "Delete StoredFile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "original_hash"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
