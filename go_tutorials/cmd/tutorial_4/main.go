package main

import (
	"fmt"
)

func main() {
	var intArray [5]int32

	fmt.Println(intArray[0])

	floatArray := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(floatArray[0])

	var intSlice []int32 = []int32{1, 2, 3, 4, 5}
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 6)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice))

	var intMap map[string]uint8 = make(map[string]uint8)
	intMap["one"] = 1
	fmt.Println(intMap["one"])

	var userMap = map[string]uint8{
		"arthur": 19,
		"john":   20,
	}

	fmt.Println(userMap["arthur"])

	delete(userMap, "arthur")
}
