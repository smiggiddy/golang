/*
Primitive data types in go
*/
package main

import (
	"fmt"
)

func main() {
	var n uint16 = 42
	fmt.Printf("%v, %T\n", n, n)

	a := 10
	b := 3

	fmt.Println(a + b)
	//go does not do inplicit conversion for you. cannot mix types

	ab := 10 //1010
	bb := 3  // 0011
	fmt.Println(ab & bb)
	fmt.Println()

	var n2 float64 = 3.14

	// go operators + - * /
	fmt.Println(n2 * 3)

	// String in go are UTF-8 by default
	// strings are inmmuatable , strings are bytes or alias for byte
	s := "this is a string ninja"
	fmt.Printf("%v, %T\n", s, s)
	beez := []byte(s) //good to work with strings and stuff and other services
	fmt.Println(beez) // converts string to bytes can be useful for pics etc

	// UTF-32 a runes char
	r := 'a' //runes use single quotes
	fmt.Printf("%v, %T\n", r, r)
}
