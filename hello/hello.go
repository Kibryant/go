package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message := greetings.Hello("Arthur")
	fmt.Println(message)
}
