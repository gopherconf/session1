package main

import (
	"fmt"
	"math"
)

/*type g interface {
	m(string)string
	n(string)string
}

type st2 struct {

}

func (s st2) m(str string)string  {
	return "Hello "+ str
}
func (s st2) n(str string)string  {
	return "Hello "+ str
}*/


type obj interface {
	area()int
}

type rect struct {
	w int
	h int
}

func (o rect)area()int{
	return o.w*o.h
}

func (o rect)width()int{
	return o.w
}

type circle struct {
	r int
}

func (o circle)area()int{
	return int(float64(o.r*o.r)*math.Pi)
}


func main(){
	/*i := st2{}

	gType(i)*/

	// anyType("ddd")

	r := rect{
		w: 100,
		h: 200,
	}

	c := circle{
		r: 100,
	}
	PrintArea(r)
	fmt.Println()
	PrintArea(c)

}

func anyType(any interface{})  {
	fmt.Printf("%+v",any)
}


func PrintArea(o obj){
	fmt.Println(o.area())
}
/*func gType(gt g)  {
	fmt.Printf("%s",gt.m("World"))
}*/