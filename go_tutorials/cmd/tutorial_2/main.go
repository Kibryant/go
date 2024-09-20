package main

import "fmt"

func main() {
	var intNumber int = 10
	fmt.Println(intNumber)

	var floatNumber float64 = 10.2
	fmt.Println(floatNumber)

	var result float64 = float64(intNumber) + floatNumber
	fmt.Println(result)

	var stringText string = "Hello, World!"
	fmt.Println(stringText)

	var boolValue bool = true
	fmt.Println(boolValue)

	const constantValue int = 10
	fmt.Println(constantValue)

	const PI float64 = 3.14159
	fmt.Println(PI)
}
