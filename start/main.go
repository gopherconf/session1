package main

import "fmt"

var i int
var arr []string

const c = "It's a const"
func main()  {
	// i := 0
	// var i int

	//str := "String"
	fmt.Printf("Hello: %s\n",c)

	arr = []string{"aaa","bbb","ccc"}
	fmt.Printf("%+v",arr)
}