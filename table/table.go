package table

import (
	"errors"
	"fmt"
	"github.com/Colk-tech/GoBase256Lib/config"
	"github.com/Colk-tech/GoBase256Lib/utils"
)

type FullWidthBytePare struct {
	Index int
	Value rune
}

type FullWidthByteTable struct {
	Contents       []FullWidthBytePare
	KeyMap         map[uint8]rune
	NullExpression rune
}

func (table FullWidthByteTable) Validate() (err error){
	nullRune, _ := utils.StrToRune("")
	lengthOfArray := len(table.Contents)

	if len(table.KeyMap) != config.NumberOfWords {
		err = errors.New("error: KeyMap is not set")
	}

	if lengthOfArray != config.NumberOfWords {
		err = fmt.Errorf(
			"error: number of words must be %d, but got %d. you can change this in `config.go`",
			config.NumberOfWords, lengthOfArray)
	}

	for i, value := range table.Contents {
		if value.Value == nullRune {
			err = fmt.Errorf("error: incorrect nil entry found at %d", i)
		}

		if value.Index != i {
			err = fmt.Errorf("error: the table is not sorted or missing key at the place %d", i)
		}
	}

	return err
}

func (table *FullWidthByteTable) CreateMap() () {
	creatingMap := map[uint8]rune{}

	for _, pare := range table.Contents {
		creatingMap[uint8(pare.Index)] = pare.Value
	}

	table.KeyMap = creatingMap
}
