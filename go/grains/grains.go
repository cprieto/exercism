package grains

import (
	"errors"
	"math/big"
)

// NOTE: You could use a bit shift operation instead of math.big
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Square out of bonds")
	}

	base := big.NewInt(2)
	return base.Exp(base, big.NewInt(int64(n)-1), nil).Uint64(), nil
}

// NOTE: Here math big is the saviour
// We are very lucky, because 2^64 - 1 is is the limit for uint64
func Total() uint64 {
	n := big.NewInt(int64(64))
	base := big.NewInt(2)
	exp := base.Exp(base, n, nil)

	return exp.Sub(exp, big.NewInt(1)).Uint64()
}
