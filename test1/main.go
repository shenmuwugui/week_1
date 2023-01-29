package main

import "fmt"

func main() {
	fb(4, 25, 16)
}

func fb(a, b, c int) {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	fmt.Printf("%v < %v < %v\n", a, b, c)
}
