package main

import "fmt"

func main() {
	//key ==>v alue
	//var v map[string]int = make(map[string]int)
	
	// one way to initialize map 
	x := make(map[string]int)
	x["first"] = 1
	x["second"] = 2
	fmt.Println(x["first"])

	//another way to initialize map
	y := map[string]int {
		"third": 3,
		"fourth": 4,
	}

	_, ok := x["third"]
	fmt.Println(ok)
	// underscore means you dont plan to use it
	// in expression

	if v, ok := x["second"]; ok {
		fmt.Println(v)
		fmt.Println("this is v")
	}

	fmt.Println(x)
	delete (x, "first")
	fmt.Println("x after deleted")
	fmt.Println(x)
	fmt.Println("this is y")
	fmt.Println(y)

	createStruct()

}

type person struct {
	Name string
	Age int
	Address string
}

func createStruct() {

	// this is structs
	// similar to classes 
	// host other data types and values

	// to create a struct:

	jason := person {
		Name: "jason",
		Age: 35,
		Address: "Germany",
	}

	fmt.Println(jason.Name)
	fmt.Println(jason)
}