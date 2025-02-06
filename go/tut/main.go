package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) sayHello() {
	fmt.Println("Hello my name is ", p.Name)

}

func main() {

	p := Person{Name: "John", Age: 25}
	fmt.Println(p)
	p.sayHello()

}
