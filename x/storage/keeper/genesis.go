package keeper

import (
	"context"

	"flstorage/x/storage/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.StoredFileMap {
		if err := k.StoredFile.Set(ctx, elem.OriginalHash, elem); err != nil {
			return err
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	if err := k.StoredFile.Walk(ctx, nil, func(_ string, val types.StoredFile) (stop bool, err error) {
		genesis.StoredFileMap = append(genesis.StoredFileMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
