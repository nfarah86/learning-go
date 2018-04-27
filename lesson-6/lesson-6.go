package main

import (
	"fmt"
	"time"
)


func main() {
	buffch := make (chan int, 5) //can store 5 int values
	buffch <- 3 // FIFO
	buffch <- 2
	//buffch <- 4

	//fmt.Println(<-buffch) // 3
	//fmt.Println(<-buffch) // 2
	
	//below statement will give an error becaue buffch doesn't have a value to print
	// buffered channels block when they are full or empty
	// here we have 2 values: 3 and 2. 
	
	//fmt.Println(<-buffch) 
	
	// ******************************************************************************
	//fmt.Println(<-waitAndSend(2,10))

	select{ // selects waits on multiple channel, until one of the channels returns whether we are listening or sending a value to a channel

		// if 2 channels are ready at the same time, go randomly selects 1
		case v1 := <- waitAndSend(5,4):
			// <- means it is listening to the channel 
			// channel generators are pop. in go;
			// create func that the outside world can listen to
			// no need to pass channels as arguments
			fmt.Println(v1)
		case v2 := <- waitAndSend(3,1):
			fmt.Println(v2)
		// prevent select statemnt from blocking code if no channels are ready
		
		// comment out default and code to see how the channels works
		default:
			fmt.Println("all channels are slow");
			// default case is evoked to prevent a select statement from blocking code if no channels are ready
	}
}

// above in the select, we have to case statements, which channel will return first?
// the one with the lowest seconds, i.

func waitAndSend(v, i1 int )chan int { // func will wait for i seconds before sending value v on return channel
	// this type of function is called a channel generator
	// which is a concurrency design pattern in go language 
	retCh := make(chan int)
	go func () {
		// creat a go routine inside the go func
		// go channel generator
		time.Sleep(time.Duration(i1)* time.Second)
			// typecast i into a time duration type that go understands
		retCh <- v // send value to the channel
		 
	}()
	return retCh // return channel to the outside world
}  // will wait for 1 second before sending value v on the return channel