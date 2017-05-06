package set1

import (
	"regexp"
	"encoding/hex"
)

func XorBufferWithChar(b []byte, c byte) []byte {
	out := []byte{}
	for i := range b {
		out = append(out, b[i]^c)
	}
	return out
}

func ScoreBufferForEnglish(b []byte) (int, error) {
	matches := 0
	for _, char := range b {
		m, err := regexp.Match("[A-Za-z ,.'\"!?()+/]", []byte{char})
		if err != nil {
			return 0.0, err
		}
		if m {
			matches++
		}
	}
	if matches == 0 {
		return 0, nil
	}

	return matches, nil
}

func GetXORScores(in []byte) ([]int, error) {
	scores := make([]int, 0)
	for i := 0; i < 256; i++ {
		b := XorBufferWithChar(in, byte(i))
		score, err := ScoreBufferForEnglish(b)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func FindIntIndexAndMax(i []int) (index, max int) {
	for i, m := range i {
		if m > max {
			index = i
			max = m
		}
	}
	return
}

func DetectXORBuffer(in []byte) ([]byte, int, error) {
	h := make([]byte, hex.DecodedLen(len(in)))
	hex.Decode(h, in)

	scores, err := GetXORScores(h)
	if err != nil {
		return nil, 0, err
	}

	index, max := FindIntIndexAndMax(scores)
	return XorBufferWithChar(h, byte(index)), max, nil
}