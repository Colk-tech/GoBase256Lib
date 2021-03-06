package engine

import (
	"github.com/Colk-tech/GoBase256Lib/utils"
)

func Decode(input string, table map[uint8]rune) (result []byte, err error) {
	reversedTable := utils.ReplaceIndexAndValue(table)

	for _, character := range input {
		b, ok := reversedTable[rune(character)]

		if !ok {
			continue
		}

		result = append(result, byte(b))
	}

	return result, err
}
