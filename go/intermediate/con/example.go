package main

import (
	"context"
	"fmt"
	"time"
)

type userKey string

const userkey userKey = "user"

func main() {
	fmt.Println("Context Example in Go")
	bgCtx := context.Background()
	fmt.Println("Background Context :", bgCtx)

	// Creating a context with a value
	ctxWithValue := context.WithValue(bgCtx, userkey, "exampleUser")
	fmt.Println("Context with Value - User : ", ctxWithValue.Value(userkey))

	// Creating a context with a timeout
	ctxWithTimeout, cancelTimeout := context.WithTimeout(bgCtx, 2*time.Second)
	defer cancelTimeout()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Context with timeout - Before Timeout : ", time.Now())
	}()

	<-ctxWithTimeout.Done()
	fmt.Println("Context with timeout - After Timeout ", time.Now())

	// Creating a context with a deadline
	deadline := time.Now().Add(5 * time.Second)
	ctxWithDeadline, cancelDeadline := context.WithDeadline(bgCtx, deadline)
	defer cancelDeadline()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Context with Deadline - ", time.Now())
	}()

	<-ctxWithDeadline.Done()
	fmt.Println("Context with Deadline - After Deadline ", time.Now())

	// Creating a context with cancellation
	ctxWithCancel, cancelFunc := context.WithCancel(bgCtx)
	defer cancelFunc()

	go func() {
		time.Sleep(2 * time.Second)
		cancelFunc() // Cancel the context
	}()

	<-ctxWithCancel.Done()
	fmt.Println("Context with cancellation - After Cancellation ", time.Now())
}

