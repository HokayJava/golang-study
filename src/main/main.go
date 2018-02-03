package main

import (
	"fmt"
	"strconv"
)

type (
	byte int8
	rune int32
)

const (
	a2 = 1
	a3
	a4
	a5 = "124"
	a6 = len(a5)
)





func main() {

	a, _, c, d := 1, 3, 4, 5
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("=====================")

	var a1 float32 = 1.1
	fmt.Println(a1)
	b := int(a1)
	fmt.Println(b)

	fmt.Println("==========数字转String===========")
	str := 5
	b1 := strconv.Itoa(str)
	fmt.Println(b1)

	fmt.Println("==========const===========")
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a6)



}
