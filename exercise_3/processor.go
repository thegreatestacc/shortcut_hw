package main

import (
	"fmt"
	"time"
)

// Обрабатывает задачи посредством чтения из канала и обычным принтом в консоль.
// Если получаем сигнал из stopCh, то программа завершает выполенение в процессе обработки задачи.
func ProcessJob(job int, stopCh <-chan struct{}, timeout time.Duration) bool {
	fmt.Printf("worker received job %d\n", job)

	select {
	case <-time.After(timeout):
		fmt.Printf("job %v timeout\n", job)
		return true
	case <-stopCh:
		fmt.Printf("worker stopped while processing a job %v\n", job)
		return false
	case <-time.After(500 * timeout):
		fmt.Printf("worker finished job %v\n", job)
		return true
	}
}
