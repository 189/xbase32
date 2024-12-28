package base32

import (
	"bytes"
	"errors"
	"strings"
)

func readChar(alphabet string, char rune) (int, error) {
	idx := strings.IndexRune(alphabet, char)
	if idx == -1 {
		return -1, errors.New("invalid character found: " + string(char))
	}
	return idx, nil
}

func Decode(input string, variant string) ([]byte, error) {
	var alphabet string

	switch variant {
	case "RFC3548", "RFC4648":
		alphabet = RFC4648
		input = strings.TrimRight(input, "=")
	case "RFC4648-HEX":
		alphabet = RFC4648_HEX
		input = strings.TrimRight(input, "=")
	case "Crockford":
		alphabet = CROCKFORD
		input = strings.ToUpper(input)
		input = strings.ReplaceAll(input, "O", "0")
		input = strings.ReplaceAll(input, "I", "1")
		input = strings.ReplaceAll(input, "L", "1")
	default:
		return nil, errors.New("unknown base32 variant: " + variant)
	}

	bits := 0
	value := 0

	output := bytes.NewBuffer(nil)

	for _, char := range input {
		charValue, err := readChar(alphabet, char)
		if err != nil {
			return nil, err
		}
		value = (value << 5) | charValue
		bits += 5

		if bits >= 8 {
			output.WriteByte(byte((value >> (bits - 8)) & 255))
			bits -= 8
		}
	}

	return output.Bytes(), nil
}
