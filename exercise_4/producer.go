package main

import (
	"context"
	"fmt"
)

// Продюсер геренирует задачи и слушает сигнал отмены из контекста.
// Если сигнал есть - отменяем выполнение, если нет - генерируем задачи и кладем в канал.
func ProduceTask(ctx context.Context, jobs chan<- int, jobsCount int) {
	fmt.Println("Producer started generate jobs...")

	defer close(jobs) // TODO discuss it

	for i := 1; i <= jobsCount; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker pool stopped: %v\n", ctx.Err())
			return
		case jobs <- i: // TODO can we write function and call it here? discuss it
			fmt.Printf("Producer added job %d\n", i)
		}
	}

	fmt.Println("Producer finished generate jobs...")
}
