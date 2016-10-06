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

	data, err := xlsx.Json()
	doError(err)

	for _, entry := range data {
		fmt.Println(string(entry))
	}
}
