package utils

import (
	crand "crypto/rand"
	"log"
	"math/big"
	"math/rand"
	// "math/rand"
)

func GenerateUnqueUrl() string {
	var randomChar = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var randomCharLength = len(randomChar)
	const uniqueStringLength = 32
	const urlLength = 10
	str := make([]rune, uniqueStringLength)
	for char := range str {
		nBig, err := crand.Int(crand.Reader, big.NewInt(int64(randomCharLength)))
		if err != nil {
			log.Printf("error occured during generating random number:%v", err)
		}
		str[char] = randomChar[nBig.Int64()]
	}

	start := rand.Intn(uniqueStringLength - urlLength)
	return string(str[start : start+urlLength])
}
