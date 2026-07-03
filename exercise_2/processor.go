package main

import (
	"fmt"
	"sync"
	"time"
)

// Логика работы процессора - прочитать число из канала jobs, возвести в квадрат, записать в канал result.
// Sleep используется для имитации длительной работы.
func ProcessJob(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d start processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)

		result := job * job

		fmt.Printf("worker %d finish processing job %v and get result %v\n", id, job, result)

		results <- result
	}
}
