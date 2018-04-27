package main 

import (
	"fmt"
)

type testiface interface {
	//define body
	// interface in go same as any langugage
	// language construct that defines functionality
	// encapsulate a lot of funcs/methods
	// can define as many methods inside an interface
	// interfaces host no more than 4 methods; usually they are small

	SayHello()
	Say(s string)
}

type testConcreteImplementation struct {}
	// implements the methods the interface supports
	// receiver syntax 

func (variable testConcreteImplementation) SayHello() {
	// (var testConcreteImplementation) is called the receiver
	// SayHello belongs to type testConcreteImplementation

	fmt.Println("Bonjour")
}

func (variable testConcreteImplementation)Say(s string) {
	fmt.Println(s, "Marsha")
}

// say and sayhello belong to testConcreteImplementation struct type


func main() {
	var tiface testiface
	tiface = testConcreteImplementation{} // called struct literal 
	tiface.SayHello()
	tiface.Say("Hello again!")
}



// interface is a type in go
// can envoke interfaces 