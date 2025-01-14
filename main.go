package main

import (
	"backdor/internal"
	"backdor/utils"
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	// Генерация простого числа P
	P := big.NewInt(rand.Int63())
	for !P.ProbablyPrime(0) {
		P = big.NewInt(rand.Int63())
	}
	// Формирования закртого ключа и параметра G
	X := big.NewInt(rand.Int63n(P.Int64()))
	G := big.NewInt(rand.Int63n(P.Int64()))
	// Класс, предоставляющий прямую и обратную функию рандомизации параметров
	randomator := utils.NewRandomator()

	// Формирование злоумышленника
	eve := internal.NewAttacker(X, P, G, randomator)
	fmt.Printf("Ключи Евы: P = %d, G = %d, X = %d, Y = %d\n", eve.PublicKey.P,
		eve.PublicKey.G, eve.PrivateKey, eve.PublicKey.Y)

	// Формирование Алисы
	user := internal.NewUser(eve.PublicKey, randomator)
	pb, _ := user.GeneratePublicKey()
	fmt.Printf("Ключи Алисы: p = %d, g = %d, y = %d, x = %d,\n", pb.P, pb.G, pb.Y, user.GetPrivateKey())

	fmt.Println("Введите сообщение для передачи")
	var m int64
	fmt.Scanf("%d", &m)

	k := big.NewInt(rand.Int63n(P.Int64()))

	// Зашифрованное сообщение
	r, s := internal.Encrypt(big.NewInt(m), pb.G, pb.Y, k, pb.P)
	fmt.Printf("Шифруется сообщение Бобом: m = %d, r = %d, s = %d\n", m, r, s)

	// Расшифровка сообщения закрытым ключом Алисы
	msg := internal.Decrypt(s, r, user.GetPrivateKey(), pb.P)
	fmt.Printf("Алиса расшифровывает сообщение Боба: m = %d\n", msg)

	// Получение закрытого ключа Алисы с помощью бэкдора
	encryptedPr := eve.DecryptUserPrivateKey(pb)
	fmt.Printf("Ева формирует закрытый ключ Алисы: x = %d\n", encryptedPr)

	// Расшифровка сообщения Евой
	encrypedMsg := internal.Decrypt(s, r, encryptedPr, pb.P)
	fmt.Printf("Ева расшифровывает сообщение Боба: m = %d\n", encrypedMsg)
}
