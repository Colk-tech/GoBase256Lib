package engine


func Encode(input []byte, table map[uint8]rune) (result string, err error) {
	for _, value := range input {
		result += string(table[value])
	}

	return result, err
}
