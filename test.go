package main

import (
	"fmt"

	"web3/internal/hash"
)

func main() {
	var pass string
	fmt.Scanln(&pass)

	passHash, err := hash.HashPassword(pass)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	fmt.Printf("hash: %v\n", passHash)

	fmt.Scanln(&pass)
	passHash2, err := hash.HashPassword(pass)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	fmt.Printf("hash: %v\n", passHash2)
	fmt.Println(hash.CheckPassword(pass, passHash))
}
