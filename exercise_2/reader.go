package main

import "fmt"

// Читает полученные значения из канала results и принтит их.
func ReadResult(results <-chan int) {
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
