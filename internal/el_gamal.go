package internal

import (
	"math/big"
)

// Шифрования Эль-Гамалем
func Encrypt(m, g, y, k, p *big.Int) (*big.Int, *big.Int) {
	r := big.NewInt(0).Exp(g, k, p)
	exp := big.NewInt(0).Exp(y, k, p)
	s := big.NewInt(0).Mod(big.NewInt(0).Mul(m, exp), p)
	return r, s
}

// Расшифрование Эль-Гамалем
func Decrypt(s, r, x, p *big.Int) *big.Int {
	msg := big.NewInt(0).Exp(r, x, p)
	msg.ModInverse(msg, p).Mul(msg, s).Mod(msg, p)
	return msg
}
