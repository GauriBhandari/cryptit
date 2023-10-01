package main

import (
	"fmt"
	"github.com/GauriBhandari/cryptit/decrypt"
	"github.com/GauriBhandari/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Golang")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
