package batching

import (
	"fmt"
	"sync"
)

func Batcher(data []string, size int, fn func(data []string)) {
	var wg sync.WaitGroup
	var batch []string
	var batchIndex int64

	for index, entry := range data {
		batch = append(batch, entry)

		if index%size == 0 {
			batchIndex++

			go func(data []string, marker int64) {
				wg.Add(1)
				fmt.Printf("Starting batch [%v]\n", marker)
				fn(data)
				wg.Done()
			}(batch, batchIndex)

			batch = nil
		}
	}
	wg.Wait()
}
