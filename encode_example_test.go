package base32

import (
	"fmt"
)

func ExampleEncode() {
	// Example 1: Encode with RFC4648
	data1 := []byte("Hello, World!")
	variant1 := "RFC4648"
	encoded1, err1 := Encode(data1, variant1, nil)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Encoded:", encoded1)
	}

	// Example 2: Encode with Crockford
	data2 := []byte("Base32Test")
	variant2 := "Crockford"
	options2 := &Options{Padding: new(bool)}
	*options2.Padding = false
	encoded2, err2 := Encode(data2, variant2, options2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Encoded:", encoded2)
	}

	// Example 3: Unknown variant
	data3 := []byte("InvalidVariant")
	variant3 := "Unknown"
	_, err3 := Encode(data3, variant3, nil)
	if err3 != nil {
		fmt.Println("Error:", err3)
	}

	// Output:
	// Encoded: NBSWY3DPEB3W64TMMQ======
	// Encoded: D1JPRV3F41VPYWKCCG
	// Error: unknown base32 variant: Unknown
}
