package main

import (
	"fmt"
	"sync"
)

const (
	numJobs    = 10
	numWorkers = 3
)

// Логика работы программы дефрагментирована на несколько функций, отвечающих за выполнение определенных задач.
func RunWorkerPool() {
	jobs := make(chan int, 5)
	results := make(chan int)
	var wg sync.WaitGroup

	for workerId := 0; workerId < numWorkers; workerId++ {
		wg.Add(1)
		go ProcessJob(workerId, jobs, results, &wg)
	}

	go GenerateWork(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	ReadResult(results)

	fmt.Println("all jobs finished")
}
