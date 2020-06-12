package diadem

import (
	"math/big"

	"github.com/diademnetwork/go-diadem/common"
	"github.com/diademnetwork/go-diadem/types"
)

type (
	BigUInt     = common.BigUInt
	BlockHeader = types.BlockHeader
)

// NewBigUint creates a biguint from a bigint
func NewBigUInt(i *big.Int) *BigUInt {
	return &BigUInt{i}
}

// NewBigUintFromInt creates a biguint from a int64
func NewBigUIntFromInt(i int64) *BigUInt {
	return &BigUInt{big.NewInt(i)}
}

func BigZeroPB() *types.BigUInt {
	return &types.BigUInt{Value: *common.BigZero()}
}
