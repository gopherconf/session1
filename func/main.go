package main

import (
	"fmt"
	"strings"
)

func main()  {
	a,b := fnName("hello World")
	fmt.Printf("a:%s - b:%s",a,b)


	multiple("\n","Gopher","Conf","...")
}

func fnName(str string) (string,string) {
	out := strings.Split(str," ")

	if len(out) >= 2{
		return out[0],out[1]
	}

	return "",""
}

func multiple(str ...string)  {
	for i := range str{
		fmt.Printf("%d - %s \n",i,str[i])
	}
}