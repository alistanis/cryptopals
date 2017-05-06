package set1

func RepeatingKeyXOR(buffer, key []byte) []byte {
	index := 0
	out := []byte{}
	for _, b := range buffer {
		if index > len(key)-1 {
			index = 0
		}
		out = append(out, b^key[index])
		index++
	}
	return out
}
