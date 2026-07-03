package main

import (
	"fmt"
	"time"
)

// Читает задачи из канала и слушает stopCh для отсановки работы.
func ReadJobs(jobs <-chan int, stopCh <-chan struct{}, timeout time.Duration) {
	for {
		select {
		case <-stopCh:
			fmt.Println("worker is done")
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("job channel closed")
				return
			}
			if canContinueProcessing := ProcessJob(job, stopCh, timeout); !canContinueProcessing {
				return
			}
		}
	}
}
