package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ProcessJobs(
	ctx context.Context,
	workerID int,
	jobs <-chan int,
	results chan<- int,
	wg *sync.WaitGroup,
	statistics *Statistics,
	semaphore *Semaphore,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped, error: %v\n", workerID, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Jobs channel is empty. Worker id is %v\n", workerID)
				return
			}

			fmt.Printf("Processing job %d\n", job)

			result, err := callWithSemaphore(ctx, job, semaphore)

			if err != nil {
				statistics.IncrementFailedTasksCounter()
				continue
			}
			select {
			case <-ctx.Done():
				return
			case results <- result:
				statistics.IncrementProcessed(workerID)
			}
		}
	}
}

func callWithSemaphore(ctx context.Context, job int, semaphore *Semaphore) (int, error) {
	err := semaphore.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer semaphore.Release()
	return callExternalService(ctx, job)
}

func callExternalService(ctx context.Context, job int) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case <-time.After(500 * time.Millisecond):
		return job, nil
	}
}
