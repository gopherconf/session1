package main

import "fmt"

type animal struct{
	Animal string
}

type gopher struct {
	animal
	Name string
	Country string
	Number int
}

func main()  {
	g := gopher{
		Name:    "Gopher",
		Country: "Iran",
		Number:  2,
	}

	fmt.Printf("%+v",g)
}
