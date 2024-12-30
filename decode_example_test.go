package xbase32

import (
	"fmt"
)

func ExampleDecode() {
	// Example 1: Valid RFC4648 decoding
	encoded1 := "JBSWY3DPFQQFO33SNRSCC==="
	variant1 := "RFC4648"
	decoded1, err1 := Decode(encoded1, variant1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Decoded:", string(decoded1))
	}

	// Example 2: Valid Crockford decoding
	encoded2 := "91IMOR3F41BMUSJCCG"
	variant2 := "Crockford"
	decoded2, err2 := Decode(encoded2, variant2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Decoded:", string(decoded2))
	}

	// Example 3: Invalid input
	encoded3 := "INVALID"
	variant3 := "RFC4648"
	_, err3 := Decode(encoded3, variant3)
	if err3 != nil {
		fmt.Println("Error:", err3)
	}

	// Output:
	// Decoded: Hello, World!
	// Decoded: Base32Test
	// Error: invalid character found: I
}
