package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a //declared as pointer

	fmt.Println(a, b)
	a = 27
	fmt.Println(a, *b)
	*b = 42
	fmt.Println(a, *b)

	// is the address operator &
	// deref operator *<var_name>
	// * before type shows a pointer to an int

	// array example
	d := [3]int{1, 2, 3}
	e := &d[0]
	c := &d[1]
	fmt.Printf("%v %p %p\n", d, e, c)

	// memory address is based on how many bites
	// go does not allow math on pointers aka pointer arithmitic
	// use unsafe package for pointer arithmitic

	// creating types
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	// every variable has an initializing value usually 0
	// pointers not intilized are nil
	var test *myStruct
	fmt.Println(test) //nil pointer
	test = new(myStruct)
	//if pointers are nil you'll get a runtime error
	(*test).foo = 33
	fmt.Println((*test).foo)

	test.foo = 22
	fmt.Println(test.foo)
	test = nil
	fmt.Println(test)
	// ^ go can handle this without needing to use the parents
	// interpreter can derefernece the pointer for us in Go
}

type myStruct struct {
	foo int
}

// types with internal pointers
// slices and maps use internal pointers
