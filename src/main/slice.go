package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1)

	a := [10]int{}
	fmt.Println(a)

	b := a[9]
	fmt.Println(b)

	c := a[5:9]
	fmt.Println(c)

	d := a[5:len(a)]
	fmt.Println(d)

	e := a[5:]
	fmt.Println(e)

	f := a[:9]
	fmt.Println(f)

	fmt.Println("====================")



}
