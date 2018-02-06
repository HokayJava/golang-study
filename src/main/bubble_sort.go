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

	j := [2] int{1, 2}
	k := [2] int{1, 2}
	fmt.Println(j == k)

	m := [10]int{}
	m[1] = 2
	fmt.Println(m)
	l := new([10]int)
	fmt.Println(l)

	fmt.Println("======================")
	n := [2][3]int{
		{1, 2, 3},
		{4, 5, 6}}
	fmt.Println(n)

	fmt.Println("======================")
	o := [...][3]int{
		{1, 2, 3},
		{4, 5, 6}}
	fmt.Println(o)

	fmt.Println("===========bubblesort===========")

	q := [...]int{1, 2, 0, 5, 3, 6}
	fmt.Println(q)
	for i := 0; i < len(q); i++ {
		for j := i + 1; j < len(q); j++ {
			if q[i] < q[j] {
				temp := q[i]
				q[i] = q[j]
				q[j] = temp
			}
		}

	}
	fmt.Println(q)

}
