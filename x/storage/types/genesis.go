package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:        DefaultParams(),
		StoredFileMap: []StoredFile{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	storedFileIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredFileMap {
		index := fmt.Sprint(elem.OriginalHash)
		if _, ok := storedFileIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedFile")
		}
		storedFileIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
