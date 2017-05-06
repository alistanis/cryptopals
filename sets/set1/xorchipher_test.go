package set1

import (
	"fmt"
	"testing"
)

func TestXorBufferWithChar(t *testing.T) {
	var testBytes = []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	b, _, err := DetectXORBuffer(testBytes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
