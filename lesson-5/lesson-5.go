package main

import (
	"fmt"
	"time"
)


func main() {
	goroutine1()
	quitSignal := make(chan bool)
	
	//anon func
	go func() {
		fmt.Println("hello from goroutine2")
		quitSignal <- true

	}() 
	go goroutine2(quitSignal)
	fmt.Println("hello from main")
	response  := <- quitSignal
	fmt.Println("the response is: ", response)

	ic := make(chan int)
	go periodicSend(ic)
	for i := range ic {
		fmt.Println(i)
	}

	//gives an error because channel is closed
		// ic <-3

		// ok is of type boolean
		// we don't need v for value, so use underscore
		_, ok := <- ic 
		fmt.Println("this is opened?: ", ok)

}

func goroutine1() {
	// goroutines: execute code with other pieces of code
	// similar to threads; except they are not real threads
	// they do not correspond to a parallel system thread
	// gorountines get scheduled on multiplex amongst the avilable os threads
	// ie. if you 4 rating system threads; then you can schedule go routines
	// the go runtime will take care the complexity

	//goroutines are light; and can be used freely
	go fmt.Println("hello from goroutine1")
	fmt.Println("hello from 1")

	// channels: are a way for different goroutines to communicate
	// you can create channels so all your stuff gets executed

}

func goroutine2(qs chan bool) {
	fmt.Println("hello from goroutine2")
	qs <- true
	
}

func periodicSend(ic chan int) {
	i := 0
	for i <= 10 {
		ic <- i
		i++
		time.Sleep(1*time.Second)
	}
	// closes this function and range fn no longer runs
	close(ic)
}