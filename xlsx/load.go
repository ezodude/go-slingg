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

func Json() error {
	var result []map[string]interface{}

	headers, err := headers()
	if err != nil {
		return err
	}

	for rowIndex, row := range loaded.Sheets[0].Rows {
		if rowIndex == 0 {
			continue
		}

		var jsonObject = make(map[string]interface{})
		for cellIndex, cell := range row.Cells {
			cellValue, _ := cell.String()
			jsonObject[headers[cellIndex]] = cellValue
		}

		result = append(result, jsonObject)
	}

	fmt.Println("headers", headers)
	fmt.Println("result", result)

	return nil
}

func headers() (keys []string, err error) {
	var result []string

	row := loaded.Sheets[0].Rows[0]
	for _, cell := range row.Cells {
		value, err := cell.String()
		if err != nil {
			return nil, err
		}
		result = append(result, strings.ToLower(value))
	}

	return result, nil
}
