package main

import "fmt"

func main() {
	if true {
		fmt.Println("GoLang conditionals")
	}

	m := make(map[string]string)

	// if value is in the map print the val
	if val, ok := m["smig"]; ok {
		fmt.Println(val)
	}

	number := 50
	guess := -5

	// or test
	if guess < number || number > guess {
		fmt.Println("This is cool")
	}
	if guess < number && number > guess {
		fmt.Println("This is cool")
	}

	if guess < 1 {
		fmt.Println("too low")
	} else if guess > 2 && guess < 50 {
		fmt.Println("just right")
	} else {
		fmt.Println("what is you doing?")
	}

	switch {
	case guess < 1:
		fmt.Println("think bigger")
		fallthrough // will fall through to the next case
	case guess >= 2:
		fmt.Println("Still not thinking big?")
	default:
		fmt.Println("You aint even trying?")

	}

	var i interface{} = 2

	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
	case float64:
		fmt.Println("i is a float64")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

// go lazily evaluates conditionals. if whatever cause the test to pass or fail
// go will exit the conditional early
