package main

import "fmt"


func main() {
	var pI * int // memory address of a value of type int
	var I int = 3
	pI = &I // pointer to the location of variable I
	
	fmt.Println("original I")
	fmt.Println(I)
	noIncrement(I)
	fmt.Println("pass by value")
	fmt.Println(I)
	increment(pI)
	fmt.Println("pass by pointer")
	fmt.Println(I)

	doDefer()

	testPanics()
	fmt.Println ("statement after panic")
}

func increment(pI *int){
	*pI++ //deferecencing 
	// we increment the variable of the pointer
	// it increments because we pass by pointer
	// to the original I

}

func noIncrement(I int){
	I++ 
	// What happens is that go makes a copy of I
	// and that copy is passed to the function
	// so the actual value will not be incremented
	// you don't have access to the original I
}

func doDefer() {
	// a way to tell code to execute only
	// when function is exiting 

	defer fmt.Println("world1 is executed second")
	defer fmt.Println("world 2 is executed first")
	defer fmt.Println("hello")
}

func testPanics(){

	defer func(){
		if err:= recover(); err != nil {
			fmt.Println(err)
			fmt.Println ("we recovered from panic")
		}
	}()
	// equiv; of throwing an exception
	panic("A panic happened")

}


