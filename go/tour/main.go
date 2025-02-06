package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func Abs() float64 {
}

func (v Vertex) Abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y + v.Y)

}

var c, python, java bool

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")

	fmt.Println(a, b)

	v := Vertex{1.0, 2.0}
	fmt.Println(v.Abs())

}
