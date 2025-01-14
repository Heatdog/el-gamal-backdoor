package utils

import "math/big"

type PublicKey struct {
	P *big.Int
	G *big.Int
	Y *big.Int
}
