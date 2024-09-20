package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32 = 42

	fmt.Printf("Value of p: %v\n", p)

	*p = i

	fmt.Printf("Value of p: %v\n", *p)

	var intArray [5]int32 = [5]int32{1, 2, 3, 4, 5}

	fmt.Printf("Value of intArray: %v\n", intArray)

	var squaredArray [5]int32 = square(&intArray)

	fmt.Printf("Value of squaredArray: %v\n", squaredArray)
}

func square(thing *[5]int32) [5]int32 {
	for i := 0; i < len(*thing); i++ {
		(*thing)[i] = (*thing)[i] * (*thing)[i]
	}

	return *thing
}
