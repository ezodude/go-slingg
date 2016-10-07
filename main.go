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

type Batch struct {
	Data   []string
	Marker int64
}

func batcher(data []string, size int, fn func(b *Batch)) {
	var wg sync.WaitGroup
	var batch []string
	var batchIndex int64

	for index, entry := range data {
		batch = append(batch, entry)

		if index%size == 0 {
			batchIndex++

			go func(b *Batch) {
				wg.Add(1)
				fn(b)
				wg.Done()
			}(&Batch{Data: batch, Marker: batchIndex})

			batch = nil
		}
	}
	wg.Wait()
}

func postWorker(b *Batch) {
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
	batcher(data, 10, postWorker)
}
