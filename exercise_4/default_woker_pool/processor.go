package default_woker_pool

import (
	"fmt"
	"sync"
)

// Процессор обрабатывает задачи простейшим инкрементом.
func ProcessJobs(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Processing job %d\n", job)
		job += 1
		results <- job
	}
}
