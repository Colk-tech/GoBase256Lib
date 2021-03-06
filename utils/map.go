package utils

func ReplaceIndexAndValue (input map[uint8]rune) (result map[rune]uint8) {
	composing := map[rune]uint8{}

	for index, value := range input {
		composing[value] = index
	}

	return composing
}
