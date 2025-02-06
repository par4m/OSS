package main

import (
	"fmt"
	"time"
)

func printNumbers(done chan bool) {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
	done <- true
}

func printLetters(done chan bool) {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(1 * time.Second)
	}
	done <- true
}

func sendData(sendOnlyChannel chan<- int) {
	sendOnlyChannel <- 10

}

func receiveData(receiveOnlyChannel <-chan int) {
	fmt.Println("Recevied ", <-receiveOnlyChannel)
}

func main() {

	fmt.Println("Unbuffered Channel example")
	unbufferedChan := make(chan string)
	go func() {
		unbufferedChan <- "Hello from param"
		fmt.Println("Send message on unbuffered channel")

	}()

	msg := <-unbufferedChan
	fmt.Println("Recevied: ", msg)

	fmt.Println("Buffered Channel Example")

	bufferedChan := make(chan string, 2)
	bufferedChan <- "Message 1"
	fmt.Println("Sent Message 1 on buffered Channel")
	bufferedChan <- "Message 2"
	fmt.Println("Sent Message 2 on bufferd channel")

	go func() {
		time.Sleep(2 * time.Second)
		bufferedChan <- "Message 3"
	}()

	msg1 := <-bufferedChan
	fmt.Println("Received: ", msg1)

	msg2 := <-bufferedChan
	fmt.Println("Received: ", msg2)

	msg3 := <-bufferedChan

	fmt.Println("Received: ", msg3)

}
