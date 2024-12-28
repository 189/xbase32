package main

import (
	"fmt"

	"gitub.com/189/base32"
)

func main() {
	str := "hello world"
	encoded, err := base32.Encode([]byte(str), "Crockford", &base32.Options{
		Padding: nil,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encoded)

	str = "hello world"
	encoded, err = base32.Encode([]byte(str), "RFC4648", &base32.Options{
		Padding: nil,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encoded)
}
