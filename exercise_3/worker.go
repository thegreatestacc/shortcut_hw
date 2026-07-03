package main

import (
	"fmt"
	"time"
)

func StartWork() {
	jobs := make(chan int)
	stopCh := make(chan struct{})

	go GenerateJobs(jobs)

	go ReadJobs(jobs, stopCh, time.Second)

	// Задаю timeout для имитации задержки отправки сигнала в stopCh.
	// Если ее убрать, то программа остановится мгновенно.
	time.Sleep(3 * time.Second)
	close(stopCh)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("program finished")
}
