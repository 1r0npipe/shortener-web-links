package generator

import "crypto/sha256"

func GenerateSHA256(link string) []byte {
	sha := sha256.New()
	sha.Write([]byte(link))
	return sha.Sum(nil)
}
