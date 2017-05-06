package set1

import (
	"fmt"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	b, err := HexToBase64([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		fmt.Println(string(b))
		t.Fatal("The expected string was not produced")
	}
}
