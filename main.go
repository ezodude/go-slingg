package main

import (
	"bytes"
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

func postWorker(b *batching.Batch) {
	fmt.Printf("Starting batch [%v]\n", b.Marker)
	for _, json := range b.Data {
		payload := bytes.NewBufferString(json)
		res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", payload)
		io.Copy(os.Stdout, res.Body)
	}
}

func main() {
	excelFileName := "sample.xlsx"
	err := xlsx.Load(excelFileName)
	doError(err)

	data, err := xlsx.Json()
	doError(err)
	batching.Batcher(data, 10, postWorker)
}
