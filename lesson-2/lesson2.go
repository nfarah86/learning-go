package main

import "fmt"

func main() {
	s := make([]int, 5)
	s[0], s[1], s[2], s[3], s[4] = 1, 2, 3, 4, 5
	fmt.Println ("Arrays in GO")
	arraysInGo()
	slicesInGo()
	slicesMore()
	getTwo(s, 1, 3)
}

func arraysInGo() {
	var nums [5]int = [5]int{1,2,3,4,5}
	fmt.Println(nums)
}

func slicesInGo() {
	//var nums[5]int = [5]int{1,2,3,4,5}
	mySlice := []int{1,2,3,4,5}

	//append can do 1 or more nums to add to array
	// myslice is dyanmic array
	mySlice = append(mySlice, 6,7,8)
	mySlice = append(mySlice,9)
	fmt.Println(mySlice)
}

func slicesMore() {
	//pattern for making arrays"
	  //declare the length of the array and type
	  // then write out the array, or slice it from 
	  // another array, or copy it

	s:= make([]int, 5)
	s[0], s[1], s[2], s[3], s[4] = 1,2,3,4,5
	fmt.Println("This is s")
	fmt.Println(s)

	s1 := s[1:3]
	fmt.Println("This is s1")
	fmt.Println(s1)
	fmt.Println("This is s1 with capacity")
	// even though we see 2 items in the array
	// it has more capacity then what's physically shown
	fmt.Println(s1[:cap(s1)])
	fmt.Println("This is s2 with copy from s")
	s2 := make([]int, 2)
	copy(s2, s[1:3])
	fmt.Println(s2)
	fmt.Println("This is the capacity of s2")
	fmt.Println(s2[:cap(s2)])
}

func getTwo(s []int, num1, num2 int) []int {
	// below not good because of memory leak;
	// garbage collector will never remove s
	//return s[num1, num2]

	//correct way: use copy
	s2 := make([]int, num2-num1)
	copy(s2, s[num1:num2])
	return s
}