package main

import (
	"bytes"
	"fmt"
	"github.com/ezodude/go-slingg/batching"
	"github.com/ezodude/go-slingg/cli"
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
	params, err := cli.Parse()
	doError(err)

	doError(xlsx.Load(params.ExcelFileName))

	data, err := xlsx.Json()
	doError(err)
	batching.Batcher(data, 10, func(data []string) {
		postWorker(*params.Url, data)
	})
}
