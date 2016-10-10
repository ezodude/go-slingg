package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/ezodude/go-slingg/batching"
	"github.com/ezodude/go-slingg/xlsx"
	"io"
	"net/http"
	"os"
)

func doError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func postWorker(url string, data []string) {
	for _, json := range data {
		payload := bytes.NewBufferString(json)
		res, _ := http.Post(url, "application/json; charset=utf-8", payload)
		io.Copy(os.Stdout, res.Body)
	}
}

func main() {
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
		doError(errors.New("No url specified!"))
	}

	if len(excelFileName) == 0 {
		doError(errors.New("No excel filename specified!"))
	}

	err := xlsx.Load(excelFileName)
	doError(err)

	data, err := xlsx.Json()
	doError(err)
	batching.Batcher(data, 10, func(data []string) {
		postWorker(*targetUrlPtr, data)
	})
}
