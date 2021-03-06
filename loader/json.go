package loader

import (
	"encoding/json"
	"github.com/Colk-tech/GoBase256Lib/table"
	"github.com/Colk-tech/GoBase256Lib/utils"
)

type JsonPare struct {
	Index int    `json:"index"`
	Word  string `json:"value"`
}

type JsonTable struct {
	Contents       []JsonPare `json:"table"`
}

func GetJSONLoader() JsonTable {
	return JsonTable{}
}

func GenerateTableFromJSONFileName(fileName string) (resultTable table.ByteItemTable, err error) {
	jsonTable := GetJSONLoader()

	err = jsonTable.LoadFromFileName(fileName)
	if err != nil{
		return resultTable, err
	}

	resultTable, err = jsonTable.GenerateTable()

	return resultTable, err
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

func (jsTable *JsonTable) GenerateTable() (resultTable table.ByteItemTable, err error) {
	var composing []table.ByteItem

	for _, item := range jsTable.Contents {
		word, err := utils.StrToRune(item.Word)
		if err != nil {
			return resultTable, err
		}

		composing = append(composing, table.ByteItem{
			Index: item.Index,
			Word:  word,
		})

	}

	resultTable.Contents = composing
	resultTable.CreateMap()

	err = resultTable.Validate()
	if err != nil {
		return resultTable, err
	}

	return resultTable, err
}
