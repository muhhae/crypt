package main

import (
	"fmt"

	"github.com/muhhae/crypt"
)

func main() {
	key := "LongLongKey"
	secret := "LongSecretContentLoremIpsumDolorSitAmet"

	fmt.Printf("Key : %s\n", key)
	fmt.Printf("Secret : %s\n\n", secret)

	encrypted, _ := crypt.SimpleSymmetricEncrypt(secret, key)
	fmt.Printf("Encrypted : %s\n", encrypted)

	decrypted, _ := crypt.SimpleSymmetricDecrypt(encrypted, key)
	fmt.Printf("Decrypted : %s\n\n", decrypted)

	keyFromFunc, _ := crypt.SimpleSymmetricDecrypt(encrypted, secret)
	fmt.Printf("Key : %s\n\n", keyFromFunc)

	decrypted, _ = crypt.SimpleSymmetricDecrypt(encrypted, keyFromFunc)
	fmt.Printf("decrypted with keyFromFunc : %s\n\n", decrypted)
}
