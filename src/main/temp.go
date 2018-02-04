package main

import (
	"fmt"
)

const (
	_MAX_COUNT = 'A'
	b          = iota
	c          = 'B'
	d          = iota
)

const (
	e = iota
)

func main() {
	fmt.Println(_MAX_COUNT)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("=====================")
	fmt.Println(1 ^ 2)
	fmt.Println(1 << 10)

	fmt.Println("=====================")
	/**
	6	0110
	11	1011
	&	0010
 	^   1101
	|	1111
	&^	0100
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 | 11)
	fmt.Println(6 &^ 11)
	x := 0
	if x > 0 && (10/x > 1) {
		fmt.Println("ok")
	}

	fmt.Println("=========if============")

	y := 10
	if y, z := 1, 2; y < 10 {
		fmt.Println(y)
		fmt.Println(z)
	}
	fmt.Println(y)

	fmt.Println("=========for============")

	a := 1
	for a <= 3 {
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)
	}

	fmt.Println("=========for1============")

	for a <= 5 {
		a++
		fmt.Println(a)
	}

	fmt.Println("=========for2============")

	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a)
	}

	fmt.Println("=========switch1============")

	a = 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("=========switch2============")

	a = 1
	switch {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("=========switch3============")

	switch l := 1; {
	case l >= 0:
		fmt.Println("l=0")
		fallthrough
	case l >= 1:
		fmt.Println("l=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("=========break============")

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 8 {
				break LABEL1
			}

		}
		fmt.Println("break ok")
	}

	fmt.Println("=========goto============")

	for {
		for i := 0; i < 10; i++ {
			if i > 8 {
				goto LABELgoto
			}

		}
	}
LABELgoto:
	fmt.Println("goto ok")

	fmt.Println("=========continue============")
LABE_CONTINUE:
	for i := 0; i < 10; i++ {
		for {
			continue LABE_CONTINUE
			fmt.Println(i)
		}

	}
	fmt.Println("CONTINUE ok")
}
