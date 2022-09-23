package main

import "fmt"

func main() {
	// Pointer
	// Adding number without pointer, 'a' is not affected
	a := 200
	b := a

	b++
	fmt.Println("value of a:",a)
	fmt.Println("value of b:",b)

	// Adding number with pointer, 'a' is affected
	c := &a
	*c++
	fmt.Println("value of a with pointer:",a)
	fmt.Println("value of c:",*c)
}