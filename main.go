package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func main() {
	// Need a form of: XLSX.utils.sheet_to_json(this.sheet);

	fmt.Println("Let's go-slingg")
	excelFileName := "sample.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows {
		for _, cell := range row.Cells {
			cellValue, _ := cell.String()
			fmt.Printf("%s\n", cellValue)
		}
	}
}
