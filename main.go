package main

import (
	"bytes"
	"fmt"
	"github.com/ezodude/go-slingg/xlsx"
	"io"
	"net/http"
	"os"
	"sync"
)

func doError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func batcher(data []string, size int, fn func(batch []string, batchMarker int64, wg *sync.WaitGroup)) {
	var wg sync.WaitGroup
	var batch []string
	var batchIndex int64

	for index, entry := range data {
		batch = append(batch, entry)

		if index%size == 0 {
			wg.Add(1)
			batchIndex++
			go fn(batch, batchIndex, &wg)
			batch = nil
		}
	}
	wg.Wait()
}

func doPost(batch []string, batchMarker int64, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting batch [%v]\n", batchMarker)
	for _, json := range batch {
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
	batcher(data, 10, doPost)
}
