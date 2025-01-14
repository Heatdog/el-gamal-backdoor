package internal

import (
	"backdor/utils"
	"fmt"
	"math/big"
	"math/rand"
)

type user struct {
	randomator *utils.Randomator

	public_key  *utils.PublicKey
	private_key *big.Int

	attacker_key *utils.PublicKey
}

func NewUser(attacker_key *utils.PublicKey, rand *utils.Randomator) *user {
	return &user{
		randomator:   rand,
		attacker_key: attacker_key,
	}
}

// Генерация параметра C2
func (user *user) generateC2(k *big.Int) *big.Int {
	c2 := big.NewInt(0).Mul(user.private_key, big.NewInt(0).Exp(user.attacker_key.Y, k, user.attacker_key.P))
	c2.Mod(c2, user.attacker_key.P)
	return user.randomator.RandomizeP(c2, user.attacker_key.P)
}

// Генерация параметра p пользователя
func (user *user) generateP(k *big.Int) (*big.Int, error) {
	for i := 0; i < 1000; i++ {
		p := user.generateC2(k)

		if p.ProbablyPrime(1000) && p.Cmp(user.private_key) == 1 {
			return p, nil
		}
	}
	return big.NewInt(0), fmt.Errorf("генерация параметра p произошла неудачно")
}

// Генерация параметра C1
func (user *user) generateC1(k *big.Int) *big.Int {
	c1 := big.NewInt(0).Exp(user.attacker_key.G, k, user.attacker_key.P)
	fmt.Printf("c1 = %d\n", c1)

	return user.randomator.RandomizeG(c1, user.attacker_key.P)
}

// Генерация параметра g пользователя
func (user *user) generateG(k, p *big.Int) (*big.Int, error) {
	for i := 0; i < 1000; i++ {
		g := user.generateC1(k)

		if g.Cmp(p) == -1 {
			return g, nil
		}
	}
	return big.NewInt(0), fmt.Errorf("генерация параметра g произошла неудачно")
}

// Генерация публичных ключей пользователя
func (user *user) GeneratePublicKey() (*utils.PublicKey, error) {
	user.private_key = big.NewInt(rand.Int63n(big.NewInt(0).Sub(user.attacker_key.P, big.NewInt(1)).Int64()))
	k := big.NewInt(rand.Int63())

	p, err := user.generateP(k)
	if err != nil {
		return nil, err
	}

	g, err := user.generateG(k, p)
	if err != nil {
		return nil, err
	}

	pb_key := &utils.PublicKey{
		P: p,
		G: g,
		Y: big.NewInt(0).Exp(g, user.private_key, p),
	}

	user.public_key = pb_key

	return pb_key, nil
}

func (u *user) GetPrivateKey() *big.Int {
	return u.private_key
}
