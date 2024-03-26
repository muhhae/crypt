package main

import (
	"encoding/base64"
	"fmt"

	"github.com/muhhae/crypt"
)

func findSecret(source, encrypted string) (string, error) {
	byteC, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	byteK := []byte(source)
	res := make([]byte, len(byteC))
	k := 0
	for i := range byteC {
		res[i] = byteC[i] ^ byteK[k]
		k++
		if k >= len(byteK) {
			k = 0
		}
	}
	return string(res), nil
}

func main() {
	key := "LongLongKey"
	secret := "LongSecretContentLoremIpsumDolorSitAmet"

	fmt.Printf("Key : %s\n", key)
	fmt.Printf("Secret : %s\n\n", secret)

	encrypted, _ := crypt.SimpleSymmetricEncrypt(secret, key)
	fmt.Printf("Encrypted : %s\n", encrypted)

	decrypted, _ := crypt.SimpleSymmetricDecrypt(encrypted, key)
	fmt.Printf("Decrypted : %s\n\n", decrypted)

	secretFromFunc, _ := findSecret(secret, encrypted)
	fmt.Printf("Key %s\n\n", secretFromFunc)

	encrypted, _ = crypt.ModifiedSimpleSymmetricEncrypt(secret, key)
	fmt.Printf("Modified Encrypted : %s\n", encrypted)

	decrypted, _ = crypt.ModifiedSimpleSymmetricDecrypt(encrypted, key)
	fmt.Printf("Modified Decrypted : %s\n\n", decrypted)

	secretFromFunc, _ = findSecret(secret, encrypted)
	fmt.Printf("Key %s\n\n", secretFromFunc)
}
