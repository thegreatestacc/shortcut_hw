package main

import (
	"fmt"
)

// Логика работы генератора - записать числа от 0 до numJobs в канал jobs.
func GenerateWork(jobs chan<- int) {
	for job := 0; job < numJobs; job++ {
		fmt.Printf("send job %d\n", job)
		jobs <- job
	}
	close(jobs)
}
