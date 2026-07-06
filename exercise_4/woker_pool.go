package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	workersCount  = 3
	jobsChanSize  = 5
	jobsCount     = 5
	semaphoreSize = 2
)

func RunWorkerPool() {

	var wg sync.WaitGroup

	// TODO looks like we can create config for all objects that we gonna generate before program is started
	jobs := make(chan int, jobsChanSize)
	results := make(chan int, jobsChanSize)

	statistics := NewStatistics()
	semaphore := NewSemaphore(semaphoreSize)

	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ctx, cancel := context.WithTimeout(signalCtx, 3*time.Second)
	defer cancel()

	go ProduceTask(ctx, jobs, jobsCount)

	fmt.Printf("Worker pool started generate workers...\n")
	for workerID := 0; workerID < workersCount; workerID++ {
		wg.Add(1)
		fmt.Printf("Worker %d starting...\n", workerID)
		go ProcessJobs(ctx, workerID, jobs, results, &wg, statistics, semaphore)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Processed result is %d \n", result)
	}

	statistics.Print()
	fmt.Printf("Worker pool finished.\n")
}
