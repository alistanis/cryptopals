package set1

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	b1 := []byte("1c0111001f010100061a024b53535009181c")
	b2 := []byte("686974207468652062756c6c277320657965")

	b1, b2, err := HexDecodeBuffers(b1, b2)
	if err != nil {
		t.Fatal(err)
	}

	b3, err := FixedXOR(b1, b2)
	if err != nil {
		t.Fatal(err)
	}

	h := make([]byte, hex.EncodedLen(len(b3)))
	hex.Encode(h, b3)
	if string(h) != "746865206b696420646f6e277420706c6179" {
		fmt.Println(string(h))
		t.Fatal("The expected output did not match")
	}
}
