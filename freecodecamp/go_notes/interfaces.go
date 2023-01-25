// interfaces in golang
// basics composing interfaces type conversion
// the empty interface and type switches and best practices
package main

import "fmt"

func main() {

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("This is an int")
	default:
		fmt.Println("No idea what type you are doc")
	}

	// i, ok = i.(float64)
	// fmt.Println(ok)
}

// interface definition
// don't describe data but defines actions
type Writer interface {
	Write([]byte) (int, error) //writing a slice of bytes to something
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*

Interface Best practices

Use many, small interfaces vs large monolithic
io.Writer, io.Reader, interface{}

don't export interfaces for types that will be consumed




*/
