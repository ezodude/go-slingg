package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Params struct {
	Url           *string
	ExcelFileName string
}

func Parse() (params *Params, err error) {
	var excelFileName string
	targetUrlPtr := flag.String("url", "", "URL to http POST to.")

	flag.Usage = func() {
		fmt.Printf("\n  Usage: %s [flags] xlsx_file\n", os.Args[0])
		fmt.Printf("\n  Slingg translates rows in an .XLSX file to JSON POST requests.\n")
		fmt.Printf("\n  Options:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	for index, arg := range flag.Args() {
		if index == 0 {
			excelFileName = arg
		}
	}

	if len(*targetUrlPtr) == 0 {
		return nil, errors.New("No url specified!")
	}

	if len(excelFileName) == 0 {
		return nil, errors.New("No excel filename specified!")
	}

	return &Params{Url: targetUrlPtr, ExcelFileName: excelFileName}, nil
}
