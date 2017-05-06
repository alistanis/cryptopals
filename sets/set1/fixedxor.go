package set1

import (
	"encoding/hex"
	"errors"
)

var (
	ErrBuffersNotEqualLength = errors.New("The buffers given were not of equal length.")
)

func HexDecodeBuffers(in1, in2 []byte) ([]byte, []byte, error) {
	b1 := make([]byte, hex.DecodedLen(len(in1)))
	b2 := make([]byte, hex.DecodedLen(len(in2)))
	_, err := hex.Decode(b1, in1)
	if err != nil {
		return nil, nil, err
	}
	_, err = hex.Decode(b2, in2)
	if err != nil {
		return nil, nil, err
	}
	return b1, b2, nil
}

func FixedXOR(b1, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, ErrBuffersNotEqualLength
	}
	b3 := []byte{}
	for i := range b1 {
		b3 = append(b3, b1[i]^b2[i])
	}
	return b3, nil
}
