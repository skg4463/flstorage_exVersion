package types_test

import (
	"testing"

	"flstorage/x/storage/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{StoredFileMap: []types.StoredFile{{OriginalHash: "0"}, {OriginalHash: "1"}}},
			valid:    true,
		}, {
			desc: "duplicated storedFile",
			genState: &types.GenesisState{
				StoredFileMap: []types.StoredFile{
					{
						OriginalHash: "0",
					},
					{
						OriginalHash: "0",
					},
				},
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
