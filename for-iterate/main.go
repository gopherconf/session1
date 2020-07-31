package main

import "fmt"

func main()  {
	slice := []string{"a","b","c","d"}

	/*for i,j := range slice{
		fmt.Printf("%d - %s \n",i,j)
	}*/

	for i := range slice{
		fmt.Printf("%d - %s \n",i,slice[i])
	}

	/*
		for i:= 0;i < 100;i++{

		}*/
}