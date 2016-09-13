package main

import (
	"fmt"
	"github.com/ezodude/go-slingg/xlsx"
	"os"
)

func main() {
	fmt.Println("Let's go-slingg")
	excelFileName := "sample.xlsx"
	err := xlsx.Load(excelFileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	xlsx.Json()
	// xlsx.Print()
}
