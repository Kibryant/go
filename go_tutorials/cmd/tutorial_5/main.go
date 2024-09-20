package main

import (
	"fmt"
	"time"
)

func main() {
	const N = 1000000
	var intSlice []int32
	var testSlice []int32 = make([]int32, 0, N)

	fmt.Printf("Time for append (uninitialized slice): %v\n", timeLoop(&intSlice, N))
	fmt.Printf("Time for append (preallocated slice): %v\n", timeLoop(&testSlice, N))
}

func timeLoop(slice *[]int32, n int) time.Duration {
	startTime := time.Now()

	for len(*slice) < n {
		*slice = append(*slice, 1)
	}

	return time.Since(startTime)
}
