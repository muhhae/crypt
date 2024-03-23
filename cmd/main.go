package main

import (
	"fmt"
	"log"

	"github.com/muhhae/crypt"
)

func main() {
	key := "LongLongKey"
	secret := "LongSecretContentLoremIpsumDolorSitAmet"

	fmt.Printf("Key : %s\n", key)
	fmt.Printf("Secret : %s\n\n", secret)

	encrypted, err := crypt.SimpleSymmetricEncrypt(secret, key)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Encrypted : %s\n", encrypted)

	decrypted, err := crypt.SimpleSymmetricDecrypt(encrypted, key)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Decrypted : %s\n\n", decrypted)

	encrypted, err = crypt.ModifiedSimpleSymmetricEncrypt(secret, key)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Modified Encrypted : %s\n", encrypted)

	decrypted, err = crypt.ModifiedSimpleSymmetricDecrypt(encrypted, key)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Modified Decrypted : %s\n\n", decrypted)
}
