package main

import "flag"
import "fmt"

func main() {
	wordPtr := flag.String("word", "foo", "this is a string")
	numPtr := flag.Int("num", 42, "this is an integer")
	boolPtr := flag.Bool ("fork", false, "this is a boolean statement")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string variable")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
    fmt.Println("num:", *numPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	
	//https://gobyexample.com/command-line-flags
}