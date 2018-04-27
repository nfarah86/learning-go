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
	Increment()
	GetInternalValue() int
}

type testConcreteImplementation struct {
	i int 
}
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

func (variable *testConcreteImplementation) Increment() {
	variable.i++
}

func (variable *testConcreteImplementation) GetInternalValue() int {
	return variable.i
}

// say and sayhello belong to testConcreteImplementation struct type

func NewTestConcreteImplementationConstructor() testiface {
	return new(testConcreteImplementation)
	// can also do &testConcreteImplementation{}
	// returns a pointer
}

func NewTestConcreteImplementationConstructorWithV(v int) testiface {
	return &testConcreteImplementation{i:v}
}

//embedding in go: encapsulate an object with another object
type testEmbedding struct { // to have all the features of testConcreteImplementation
	*testConcreteImplementation //embedding
	// able to access all the methods of testConcreteImplementation with this type
}

func testEIface(v interface{}) {
	// this is an empty interface
	// there are no methods to satisfy
	// any datatype can go through

	/*
	RUN THIS AS AN OPTION as a type assertion
	// executes if v is only int;
	if i, ok := v.(int); ok {
		fmt.Println ("I am an integer of value", i)
	} else {
		fmt.Println ("I am not an integer type")
	}
	*/

	// RUN THIS AS AN OPTION as a type switch statement

	switch val := v.(type) {
	case int:
		fmt.Println ("I'm an int ", val)
	case string:
		fmt.Println("I'm a string",val)
	default:
		fmt.Println("I'm not a string or int", val)
	}
}

func main() {
	var tiface testiface
	
	// *************** one way to write the pointer is here: ***************
	// tiface = &testConcreteImplementation{} // called struct literal 
	// & memory address of value, which is a reference
	// can also used new(testConcreteImplementation) istead of &
	
	// *************** another way to write the pointer is here: ***************
	// tiface = NewTestConcreteImplementationConstructor()


	// *************** another way to write the pointer is here: ***************
	tiface = NewTestConcreteImplementationConstructorWithV(5)
	tiface.SayHello()
	tiface.Say("Hello again!")
	tiface.Increment()
	tiface.Increment()
	fmt.Println(tiface.GetInternalValue())
	// when implemented without the * and &, you get 0. Why? it's not a pointer. The types of the receivers are not a pointer.
	// The original object doesn't retain changes that methods do to it
	// to fix, use pointer -> get passed by reference instead of value.
	// used heavily; can pass any value type, including structs, interfaces
	fmt.Println ( "Embedding example")
	te :=testEmbedding{testConcreteImplementation: &testConcreteImplementation{i:50}}
	te.SayHello()
	te.Say("Hola")
	te.Increment()
	fmt.Println(te.GetInternalValue())

	fmt.Println("********** empty interface **********")
	testEIface(3) // no error can pass any object
	testEIface("this is a string")
	testEIface(tiface)

	// if you want to specify the datatypes passed in the interface
	// can use type assertion
}

// interface is a type in go
// can envoke interfaces 