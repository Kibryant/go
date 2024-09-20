package main

import (
	"errors"
	"fmt"
)

func main() {
	var result int
	var err error

	result, err = intDivision(10, 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func intDivision(a int, b int) (int, error) {
	var err error

	if b == 0 {
		err = errors.New("division by zero")

		return 0, err
	}

	return a / b, err
}
