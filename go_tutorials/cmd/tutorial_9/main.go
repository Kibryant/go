package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var WAIT_GROUP sync.WaitGroup = sync.WaitGroup{}
var MOCK_DB_DATA = []string{"id-1", "id-2", "id-3", "id-4", "id-5"}
var RESULT = []string{}
var MUTEX sync.Mutex = sync.Mutex{}

func main() {
	startTime := time.Now()

	for i := 0; i < len(MOCK_DB_DATA); i++ {
		WAIT_GROUP.Add(1)
		go getMockData(i)
	}

	WAIT_GROUP.Wait()
	elapsedTime := time.Since(startTime)

	fmt.Println("The elapsed time is:", elapsedTime)
	fmt.Println("The result of the main function is:", RESULT)
}

func getMockData(idx int) {
	var DELAY float32 = rand.Float32() * 2000

	time.Sleep(time.Duration(DELAY) * time.Millisecond)

	fmt.Println("The result of the getMockData function is:", MOCK_DB_DATA[idx])

	saveMockData(MOCK_DB_DATA[idx])

	WAIT_GROUP.Done()
}

func saveMockData(data string) {
	MUTEX.Lock()
	RESULT = append(RESULT, data)
	MUTEX.Unlock()
}
