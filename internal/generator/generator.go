package generator

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

var ErrNoEncodePossible = errors.New("not possible to encode!")

func applySha256(link string) []byte {
	sha := sha256.New()
	sha.Write([]byte(link))
	return sha.Sum(nil)
}

func base58Enc(input []byte) string {
	encode := base58.BitcoinEncoding
	encoded, err := encode.Encode(input)
	if err != nil {
		return fmt.Sprint("error")
	}
	return string(encoded)
}

func GenerateShortUrl(input string) (string, error) {
	urlHashB := applySha256(input)
	generatedLine := new(big.Int).SetBytes(urlHashB).Uint64()
	result := base58Enc([]byte(fmt.Sprintf("%d", generatedLine)))
	if result == "error" {
		return "", ErrNoEncodePossible
	}
	fmt.Println(result)
	return result[:8], nil

}
