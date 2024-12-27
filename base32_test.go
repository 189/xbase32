package base32

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

type TestItem struct {
	input   string
	variant string
	want    []byte
	wantErr error
}

func TestDecode(t *testing.T) {
	tests := []TestItem{
		{
			input:   "MZXW6YTBOI======",
			variant: "RFC4648",
			want:    []byte("foobar"),
			wantErr: nil,
		},
		{
			input:   "CO======",
			variant: "RFC4648-HEX",
			want:    []byte{0x5c},
			wantErr: nil,
		},
		{
			input:   "91JPRV3F41BMFWPCCG",
			variant: "Crockford",
			want:    []byte("hello world"),
			wantErr: nil,
		},
		{
			input:   "MZXW6===",
			variant: "RFC4648",
			want:    nil,
			wantErr: errors.New("invalid character found: ="),
		},
		{
			input:   "INVALID",
			variant: "Unknown",
			want:    nil,
			wantErr: errors.New("unknown base32 variant: Unknown"),
		},
	}

	for _, item := range tests {
		name := fmt.Sprintf("%s_%s", item.input, item.variant)

		t.Run(name, func(t *testing.T) {
			got, err := Decode(item.input, item.variant)
			if !bytes.Equal(got, item.want) {
				t.Errorf("Decode(%q, %q) = %v, want %v", item.input, item.variant, got, item.want)
			}

			if (err == nil && item.wantErr != nil) || (err != nil && item.wantErr == nil) || (err != nil && item.wantErr != nil && item.wantErr.Error() != err.Error()) {
				t.Errorf("Decode(%q, %q) error = %v, wantErr %v", item.input, item.variant, err, item.wantErr)
			}

		})
	}

}
