package main

import "fmt"

func main() {
	var a [2] int
	var b [2] int
	b = a
	fmt.Println(b)

	c := [20] int{19: 1}
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	e := [...]int{0: 1}
	fmt.Println(e)

	f := [...]int{99: 1}
	var p *[100]int = &f
	fmt.Println(p)

	g, h := 1, 2
	i := [...]*int{&g, &h}
	fmt.Println(i)
}
