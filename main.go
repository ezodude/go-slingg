package main

import (
	"fmt"
	"github.com/ezodude/go-slingg/xlsx"
	"os"
)

func main() {
	excelFileName := "sample.xlsx"
	err := xlsx.Load(excelFileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	json, err := xlsx.Json()
	fmt.Printf("%s\n", json)
}
