package xlsx

import (
	"fmt"
	txlsx "github.com/tealeg/xlsx"
	"strings"
)

var loaded *txlsx.File
var err error

func Load(filename string) (err error) {
	loaded, err = txlsx.OpenFile(filename)
	return err
}

func Print() {
	sheet := loaded.Sheets[0]
	for _, row := range sheet.Rows {
		for _, cell := range row.Cells {
			cellValue, _ := cell.String()
			fmt.Printf("%s\n", cellValue)
		}
	}
}

func Json() {
	var result = make(map[string]interface{})

	// for _, row := range loaded.Sheets[0].Rows {
	// 	for _, cell := range row.Cells {
	// 		cellValue, _ := cell.String()

	// 	}
	// }
	fmt.Println("headers", headers())
	fmt.Println("result", result)
}

func headers() []string {
	var result []string

	row := loaded.Sheets[0].Rows[0]
	for _, cell := range row.Cells {
		value, _ := cell.String()
		result = append(result, strings.ToLower(value))
	}

	return result
}
