package xlsx

import (
	"fmt"
	txlsx "github.com/tealeg/xlsx"
)

var loaded *txlsx.File
var err error

func Load(filename string) (err error) {
	loaded, err = txlsx.OpenFile(filename)
	return err
}

func Dump() {
	sheet := loaded.Sheets[0]
	for _, row := range sheet.Rows {
		for _, cell := range row.Cells {
			cellValue, _ := cell.String()
			fmt.Printf("%s\n", cellValue)
		}
	}
}
