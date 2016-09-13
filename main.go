package main

import (
	"fmt"
	"github.com/ezodude/go-slingg/xlsx"
	"os"
)

func doError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	excelFileName := "sample.xlsx"
	err := xlsx.Load(excelFileName)
	doError(err)

	json, err := xlsx.Json()
	doError(err)
	fmt.Printf("%s\n", json)
}
