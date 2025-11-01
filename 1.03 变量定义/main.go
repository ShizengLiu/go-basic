package main

import (
	"fmt"
)

var (
	truth    = true
	nonTruth = false
)

var a, b, c string = "a", "b", "c"

func main() {
	fmt.Println("main method invoked!")
	var i int = 42
	fmt.Println(i)
	var f float64
	fmt.Println(f)
	var s = "string1"
	fmt.Println(s)
	nonVar := "nonVar"
	fmt.Println(nonVar)

}
