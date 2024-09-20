package main

import (
	"fmt"
	"time"
)

func main() {
	var channel chan int = make(chan int)
	go processChannel(channel)

	for i := range channel {
		fmt.Println("The value of i is:", i)
		time.Sleep(1 * time.Second)
	}
}

func processChannel(channel chan int) {
	defer close(channel)

	for i := 0; i < 5; i++ {
		channel <- i
	}

	fmt.Println("The processChannel function is done.")
}
