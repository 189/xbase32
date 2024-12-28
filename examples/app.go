package main

import (
	"fmt"

	base32 "gitub.com/189/xbase32"
)

func main() {
	variant := "Crockford"
	str := "hello world"

	encoded, err := base32.Encode([]byte(str), variant, &base32.Options{
		Padding: nil,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encoded) // output: D1JPRV3F41VPYWKCCG

	ret, e := base32.Decode(encoded, variant)
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(string(ret) == str) // output: true
}
