package base32

import (
	"errors"
	"strings"
)

const (
	RFC4648     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	RFC4648_HEX = "0123456789ABCDEFGHIJKLMNOPQRSTUV"
	CROCKFORD   = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
)

type Options struct {
	Padding *bool
}

func Encode(data []byte, variant string, options *Options) (string, error) {
	var alphabet string
	var defaultPadding bool

	switch variant {
	case "RFC3548", "RFC4648":
		alphabet = RFC4648
		defaultPadding = true
	case "RFC4648-HEX":
		alphabet = RFC4648_HEX
		defaultPadding = true
	case "Crockford":
		alphabet = CROCKFORD
		defaultPadding = false
	default:
		return "", errors.New("unknown base32 variant: " + variant)
	}

	padding := defaultPadding
	if options != nil && options.Padding != nil {
		padding = *options.Padding
	}

	var (
		bits   int
		value  int
		output strings.Builder
	)

	for _, b := range data {
		value = (value << 8) | int(b)
		bits += 8

		for bits >= 5 {
			output.WriteByte(alphabet[(value>>(bits-5))&31])
			bits -= 5
		}
	}

	if bits > 0 {
		output.WriteByte(alphabet[(value<<(5-bits))&31])
	}

	if padding {
		for output.Len()%8 != 0 {
			output.WriteByte('=')
		}
	}

	return output.String(), nil
}
