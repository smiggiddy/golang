package main

import "fmt"

func main() {
	g := greeter{
		greeting: "hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println(g.name)
	a := bigNumber(500)

	a.huge()
	fmt.Println(a)
}

type greeter struct {
	greeting string
	name     string
}

func (d *greeter) greet() {
	fmt.Println(d.greeting, d.name)
	d.name = "Smig"
}

type bigNumber int

func (num *bigNumber) huge() {
	fmt.Println(*num)
	*num = bigNumber(*num) * bigNumber(*num)
}
