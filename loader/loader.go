package loader

import (
	"github.com/Colk-tech/GoBase256Lib/table"
)

type Loader interface {
	LoadFromFileName(fileName string) (err error)
	LoadFromString(query string) (err error)
	GenerateTable() (resultTable table.ByteItemTable, err error)
}
