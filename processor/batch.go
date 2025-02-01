package processor

import (
	"log"
	"sync"
)

func WorkerPool(data []string, batchSize int, workerCount int) {
	var wg sync.WaitGroup
	dataChannel := make(chan []string, len(data)/batchSize)

	go func() {
		for i := 0; i < len(data); i += batchSize {
			end := i + batchSize
			if end > len(data) {
				end = len(data)
			}
			dataChannel <- data[i:end]
		}
		close(dataChannel)
	}()


	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, dataChannel)
	}

	wg.Wait()
}


func worker(wg *sync.WaitGroup, dataChannel <-chan []string) {
	defer wg.Done()
	for batch := range dataChannel {
		log.Printf("[Worker] Procesando lote de %d registros\n", len(batch))

		/*if err := db.SaveBatch(batch); err != nil {
			log.Printf("[Error al guardar en lotes] %s: ", err)
		}*/
	}
}
