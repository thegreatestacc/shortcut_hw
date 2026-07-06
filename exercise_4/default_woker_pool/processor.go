package default_woker_pool

import (
	"fmt"
	"sync"
	"time"
)

// Процессор обрабатывает задачи простейшим инкрементом.
func ProcessJobs(
	workerID int,
	jobs <-chan int,
	results chan<- int,
	wg *sync.WaitGroup,
	statistics *Statistics,
	semaphore *Semaphore,
) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Processing job %d\n", job)

		// добавлен семафор. инкрементируем счетчик
		semaphore.Acquire()
		// имитация обработки задачи
		result := callExternalService(job)
		// после обработки задачи, уменьшаем счетчик семафора
		semaphore.Release()
		results <- result

		// добавлен подсчет статистики по второй задаче
		statistics.IncrementProcessed(workerID)
	}
}

func callExternalService(job int) int {
	time.Sleep(500 * time.Millisecond)
	return job + 1
}
