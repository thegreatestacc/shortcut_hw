package default_woker_pool

import (
	"fmt"
	"sync"
)

const (
	workersCount  = 3
	jobsChanSize  = 5
	jobsCount     = 5
	semaphoreSize = 2
)

func RunWorkerPool() {
	var wg sync.WaitGroup
	jobs := make(chan int, jobsChanSize)
	results := make(chan int, jobsChanSize)

	statistics := NewStatistics()
	semaphore := make(chan struct{}, semaphoreSize)

	go ProduceTask(jobs, jobsCount)

	fmt.Printf("Worker pool started generate workers...\n")
	for workerID := 0; workerID < workersCount; workerID++ {
		wg.Add(1)
		fmt.Printf("Worker %d starting...\n", workerID)
		go ProcessJobs(workerID, jobs, results, &wg, statistics, semaphore)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Result is %d \n", result)
	}

	statistics.Print()
	fmt.Printf("Worker pool finished.\n")
}
