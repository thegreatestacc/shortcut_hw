package main

import (
	"sync"
	"testing"
)

func TestProcessJob(t *testing.T) {
	t.Run("worker must get square of a number and write it into the result channel", func(t *testing.T) {
		jobs := make(chan int, 1)
		results := make(chan int, 1)

		var wg sync.WaitGroup

		jobs <- 5
		close(jobs)

		wg.Add(1)
		go ProcessJob(1, jobs, results, &wg)

		wg.Wait()
		close(results)

		got := <-results
		want := 25

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
