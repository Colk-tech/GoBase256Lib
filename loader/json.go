package loader

import (
	"encoding/json"
	"github.com/Colk-tech/GoBase256Lib/table"
	"github.com/Colk-tech/GoBase256Lib/utils"
)

type JsonPare struct {
	Index int    `json:"Index"`
	Value string `json:"Value"`
}

type JsonTable struct {
	Contents       []JsonPare `json:"table"`
	NullExpression string     `json:"null"`
}

func (jsTable *JsonTable) LoadFromString(jsonStr string) (err error) {
	err = json.Unmarshal([]byte(jsonStr), jsTable)
	return err
}

func (jsTable *JsonTable) LoadFromFileName(fileName string) (err error) {
	jsonText, err := utils.ReadFile(fileName)

	if err != nil {
		return err
	}

	err = jsTable.LoadFromString(jsonText)

	return err
}

func (jsTable *JsonTable) ToTable() (resultTable table.FullWidthByteTable, err error) {
	resultTable.NullExpression, err = utils.StrToRune(jsTable.NullExpression)
	if err != nil {
		return resultTable, err
	}

	var composing []table.FullWidthBytePare

	for _, value := range jsTable.Contents {
		val, err := utils.StrToRune(value.Value)
		if err != nil {
			return resultTable, err
		}

		composing = append(composing, table.FullWidthBytePare{
			Index: value.Index,
			Value: val,
		})

	}

	resultTable.Contents = composing

	if err != nil {
		return resultTable, err
	}

	return resultTable, err
}
