package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Person
	JobTitle string
	Salary   float64
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) HaveBirthday() {
	p.Age++

}

func (e Employee) DisplayInfo() {
	fmt.Printf("Name: %s, Job Title: %s , Age: %d  Salary: %.2f", e.FullName(), e.JobTitle, e.Age, e.Salary)
}

func modifyValue(val *int) {
	*val = 100
}

func main() {

	x := 10
	fmt.Println("Original value", x)

	// creating a ptr to variable x
	y := &x
	fmt.Println("Address of x ", y)

	fmt.Println("Value of x", *y)

	// modified value of x
	*y = 12

	modifyValue(y)
	fmt.Println(*y)

	// go to statement
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto EndLoop
			}
			fmt.Println("i=%d, j = %d\n", i, j)
		}
	}
EndLoop:
	fmt.Println("Exited from the Nested Loop")

	// Time Package
	explainTimeMethod()

}

func explainTimeMethod() {
	t := time.Date(2024, time.June, 26, 12, 0, 0, 0, time.UTC)

	duurationToAdd := 24 * time.Hour
	newTime := t.Add(duurationToAdd)

	fmt.Println("Original Time: ", newTime)

	t2 := time.Date(2024, time.June, 27, 12, 0, 0, 0, time.UTC)
	fmt.Println("New Time: ", t2)

	fmt.Println("Starting a Ticker")
	ticker := time.Tick(1 + time.Second)
	for i := 0; i < 5; i++ {
		<-ticker
		fmt.Println("tick ", i+1)
	}

	fmt.Println("Sleeping for 2 seconds")
	time.Sleep(2 * time.Second)
	fmt.Println("Awake after sleep")

}
