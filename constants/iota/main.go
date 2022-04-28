package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)

	fmt.Println(c11, c22, c33)

	const (
		North = iota
		East
		South
		West
	)

	fmt.Println(North, East, South, West)

	const (
		a = iota * 2
		b
		c
		d
	)

	fmt.Println(a, b, c, d)

	const (
		x = -(iota + 2)
		_
		y
		z
	)

	fmt.Println(x, y, z)
}
