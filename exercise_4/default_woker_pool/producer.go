package default_woker_pool

import "fmt"

// Продюсер геренирует задачи и кладет их в канал
func ProduceTask(jobs chan<- int, jobsCount int) {
	fmt.Println("Producer started generate jobs...")
	for i := 1; i <= jobsCount; i++ {
		fmt.Printf("Producer add job %d\n", i)
		jobs <- i
	}
	close(jobs)
	fmt.Println("Producer finished generate jobs...")
}
