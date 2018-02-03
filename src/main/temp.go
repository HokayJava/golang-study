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
}
