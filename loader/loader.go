package loader

import (
	"github.com/Colk-tech/GoBase256Lib/table"
)

type Loader interface {
	LoadFromFileName(fileName string) (err error)
	LoadFromString(query string) (err error)
	ToTable() (resultTable table.FullWidthByteTable, err error)
}
