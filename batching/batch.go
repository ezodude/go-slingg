package batching

import (
	"sync"
)

type Batch struct {
	Data   []string
	Marker int64
}

func Batcher(data []string, size int, fn func(b *Batch)) {
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
