package main

import (
	"reflect"
	"testing"
)

func TestGenerateWork(t *testing.T) {
	t.Run("send all tasks and close channel", func(t *testing.T) {
		jobs := make(chan int, numJobs)

		GenerateWork(jobs)

		var got []int

		for job := range jobs {
			got = append(got, job)
		}

		want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
