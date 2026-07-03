package main

import "fmt"

// Генерирует задачи и пишет их в канал.
func GenerateJobs(jobs chan<- int) {
	for job := 0; job < 5; job++ {
		fmt.Printf("job %v has been generated \n", job)
		jobs <- job
	}
	close(jobs)
}
