package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sun(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{1, 2, 3, 4, 5, 5, 6}
	c := make(chan int)

	go sun(s[:len(s)/2], c)
	go sun(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	// go say("world")
	// say("hello")
	//
	//
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	close(ch)

	for i := range ch {
		fmt.Println(i)
	}

}
