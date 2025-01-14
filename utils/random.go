package utils

import (
	"math/big"
	"math/rand"
)

type Randomator struct {
	keyP *big.Int
	keyG *big.Int
}

// Класс, предоставляющий функции рандомизации R1 и R2
func NewRandomator() *Randomator {
	return &Randomator{}
}

func (r *Randomator) RandomizeP(val, p *big.Int) *big.Int {
	r.keyP = big.NewInt(rand.Int63n(p.Int64()))
	return big.NewInt(0).Xor(val, r.keyP)
}

func (r *Randomator) RestoreP(val, p *big.Int) *big.Int {
	return big.NewInt(0).Xor(val, r.keyP)
}

func (r *Randomator) RandomizeG(val, p *big.Int) *big.Int {
	r.keyG = big.NewInt(rand.Int63n(p.Int64()))
	return big.NewInt(0).Xor(val, r.keyG)
}

func (r *Randomator) RestoreG(val, p *big.Int) *big.Int {
	return big.NewInt(0).Xor(val, r.keyG)
}
