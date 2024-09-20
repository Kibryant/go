package main

import "fmt"

type User interface {
	GetName() string
}

type Customar struct {
	Name string
	Age  uint8
}

type Admin struct {
	Name string
	Age  uint8
}

func (c Customar) GetName() string {
	return c.Name
}

func (a Admin) GetName() string {
	return a.Name
}

func formatedPrint(u User) {
	fmt.Printf("Name: %s\n", u.GetName())
}

func main() {
	var customar Customar = Customar{
		Name: "Arthur",
		Age:  19,
	}

	var customer2 Customar = Customar{
		Name: "John",
		Age:  20,
	}

	var admin Admin = Admin{
		Name: "Admin",
		Age:  30,
	}

	formatedPrint(customar)
	formatedPrint(customer2)
	formatedPrint(admin)
}
