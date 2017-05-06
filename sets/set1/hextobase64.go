package set1

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(in []byte) ([]byte, error) {
	h := make([]byte, hex.DecodedLen(len(in)))
	_, err := hex.Decode(h, in)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer([]byte{})
	enc := base64.NewEncoder(base64.StdEncoding, b)
	_, err = enc.Write(h)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
