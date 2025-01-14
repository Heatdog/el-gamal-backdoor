package tests

import (
	"backdor/internal"
	"backdor/utils"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	testTable := []struct {
		name string

		evePubicKey   *utils.PublicKey
		evePrivateKey *big.Int

		randomator *utils.Randomator

		m int64
		k int64
	}{
		{
			name: "Пример из документа",
			evePubicKey: &utils.PublicKey{
				P: big.NewInt(23993),
				G: big.NewInt(15765),
			},
			evePrivateKey: big.NewInt(9237),

			randomator: utils.NewRandomator(),

			m: 809,
			k: 1487,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			eve := internal.NewAttacker(testCase.evePrivateKey, testCase.evePubicKey.P,
				testCase.evePubicKey.G, testCase.randomator)
			fmt.Printf("Ключи Евы: P = %d, G = %d, X = %d, Y = %d\n", eve.PublicKey.P,
				eve.PublicKey.G, eve.PrivateKey, eve.PublicKey.Y)

			user := internal.NewUser(eve.PublicKey, testCase.randomator)
			pb, _ := user.GeneratePublicKey()
			fmt.Printf("Ключи Алисы: p = %d, g = %d, y = %d, x = %d,\n", pb.P, pb.G, pb.Y, user.GetPrivateKey())

			r, s := internal.Encrypt(big.NewInt(testCase.m), pb.G, pb.Y, big.NewInt(testCase.k), pb.P)
			fmt.Printf("Шифруется сообщение Бобом: m = %d, r = %d, s = %d\n", testCase.m, r, s)

			msg := internal.Decrypt(s, r, user.GetPrivateKey(), pb.P)
			fmt.Printf("Алиса расшифровывает сообщение Боба: m = %d\n", msg)

			require.Equal(t, testCase.m, msg.Int64())

			encryptedPr := eve.DecryptUserPrivateKey(pb)
			fmt.Printf("Ева формирует закрытый ключ Алисы: x = %d\n", encryptedPr)

			require.Equal(t, user.GetPrivateKey().Int64(), encryptedPr.Int64())
		})
	}
}
