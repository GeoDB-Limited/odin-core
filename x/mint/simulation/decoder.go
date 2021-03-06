package simulation

import (
	"bytes"
	"fmt"

	minttypes "github.com/GeoDB-Limited/odin-core/x/mint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

// NewDecodeStore returns a decoder function closure that umarshals the KVPair's
// Value to the corresponding mint type.
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key, minttypes.MinterKey):
			var minterA, minterB minttypes.Minter
			cdc.MustUnmarshalBinaryBare(kvA.Value, &minterA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &minterB)
			return fmt.Sprintf("%v\n%v", minterA, minterB)
		default:
			panic(fmt.Sprintf("invalid mint key %X", kvA.Key))
		}
	}
}
