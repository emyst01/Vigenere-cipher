package main

import (
	"fmt"
	"strings"
)

func encrypt(plaintext string, key string) string {
	plaintext = strings.ToUpper(plaintext)
	key = strings.ToUpper(key)

	ciphertext := ""
	j := 0
	for i := 0; i < len(plaintext); i++ {
		if plaintext[i] < 'A' || plaintext[i] > 'Z' {
			ciphertext += string(plaintext[i])
			continue
		}

		ciphertext += string((plaintext[i]+key[j]-2*'A')%26 + 'A')
		j = (j + 1) % len(key)
	}
	return ciphertext
}

func decrypt(ciphertext string, key string) string {
	ciphertext = strings.ToUpper(ciphertext)
	key = strings.ToUpper(key)

	plaintext := ""
	j := 0
	for i := 0; i < len(ciphertext); i++ {
		if ciphertext[i] < 'A' || ciphertext[i] > 'Z' {
			plaintext += string(ciphertext[i])
			continue
		}

		plaintext += string((ciphertext[i]-key[j]+26)%26 + 'A')
		j = (j + 1) % len(key)
	}
	return plaintext
}

func main() {
	var plaintext string
	var key string
	var enc_dec int
	for true {
		fmt.Println("Type the message (only letters, no spaces and numbers)")
		fmt.Scanln(&plaintext)
		fmt.Println("Now type the key")
		fmt.Scanln(&key)
		fmt.Println("0: Encrypt\n1: Decrypt")
		fmt.Scanln(&enc_dec)
		if enc_dec == 0 {
			fmt.Println(encrypt(plaintext, key))
		} else if enc_dec == 1 {
			fmt.Println(decrypt(plaintext, key))
		} else {
			fmt.Println("The options are 0 and 1.")
		}
	}

}
