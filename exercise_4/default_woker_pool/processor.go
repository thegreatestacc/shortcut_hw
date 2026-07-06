package default_woker_pool

import (
	"fmt"
	"sync"
)

// Процессор обрабатывает задачи простейшим инкрементом.
func ProcessJobs(workerID int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup, statistics *Statistics) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Processing job %d\n", job)
		job += 1
		results <- job

		// добавлен подсчет статистики по второй задаче
		statistics.IncrementProcessed(workerID)
	}
}
