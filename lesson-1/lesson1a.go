package main

import "fmt"


func main() {
	// fn's calls
	fmt.Println("Comapre numbers")
	fmt.Println(ComapreNumbers(1,2))
	fmt.Println("Case Statement 1")
	fmt.Println(caseStatements(2,4))
	fmt.Println("Case Statement 2")
	fmt.Println(caseStatement2(4))
	fmt.Println("For Loops")
	forLoops()
	fmt.Println("While Loops")
	whileLoops()
	
	// different dec. of variables
	//x:= 2 
		//inference declaration of variable
	//var x unit8 = 2 
		//explicit declaration of variable
}

func ComapreNumbers(i1, i2 int)(bool, int) {
	if i1 > i2 {
		fmt.Println("i1 is greater than the i2")
		return false, i1-i2
	} else if (i2 > i1){
		fmt.Println("i2 is greater than i1")
		return false, i2-i1
	}
	return true, 0
}

func caseStatements (i1, i2 int)(bool, int) {
switch {		
	case i1 > i2:
		fmt.Println("i1 > i2")
		return false, i1 - i2
	case i2 > i1:
		fmt.Println ("i2 > i1")
		return false, i2- i1
	}
	return true, 0

}
func  caseStatement2(i1 int)(int) {
	switch i1 {
	case 3:
		fmt.Println("x is 3") 
	case 4:
		fmt.Println("x is 4")
		fallthrough
	case 5:
		fmt.Println ("x is 5")
	}
	return i1
}

func forLoops(){
	for i1 := 0; i1 <= 10; i1++ {
		fmt.Println(i1)
	}
}

func whileLoops(){
	i1 := 10
	for i1 >= 0 {
		fmt.Println(i1)
		i1--
	}
}
