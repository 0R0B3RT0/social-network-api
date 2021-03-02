package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	key := make([]byte, 64)

	if _, error := rand.Read(key); error != nil {
		log.Fatal(error)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
}
