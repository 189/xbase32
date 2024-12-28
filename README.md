# xbase32

Package base32 implements base32 encoding as specified by RFC 4648, RFC 4648 Hex, Crockford.

> The standard library in Golang does not support encoding/decoding as specified by Crockford

## Installation

```
$ go get -u github.com/189/xbase32
```

## Usage

```
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


```
