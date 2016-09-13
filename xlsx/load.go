package xlsx

import (
	"encoding/json"
	"fmt"
	txlsx "github.com/tealeg/xlsx"
	"strings"
)

var loaded *txlsx.File

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

func Json() (data []byte, err error) {
	var base []map[string]interface{}

	headers, err := headers()
	if err != nil {
		return nil, err
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

		base = append(base, jsonObject)
	}

	return json.Marshal(base)
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
