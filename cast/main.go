package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var i int
	i = 2
	fmt.Print(mu64(int64(i)))
	fmt.Print(mu(float32(i)))


	str := "1"
	strconv.ParseInt(str,10,32)

}

func mu(i float32) float32{
	return i*2
}

func mu64(i int64) int64{
	return i*2
}