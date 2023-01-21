package main

import (
	"fmt"
)

func main() {
	fmt.Println("This lesson is on printing")
	sayItWithYourChest("You are doing this shit!")
	twoTypes("hello", "smig")

	sum := sumVariaticPointer(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("The sum is", *sum)

	fmt.Println(test((2)))
	d, err := divide(3.4, 0.0)
	if err != nil {
		fmt.Println(err)
		// return (will end program)
	}
	fmt.Println(d)

	// function can be treated as variables, objects etc
	//anonomous function

	func() {
		fmt.Println("Inside this function")
	}()
	// function as a variable
	var t func() = func() {
		fmt.Println("Inside this function")
	}
	t()

	tT := func() {
		fmt.Println("Inside this function")
	}
	tT()
}

func sayItWithYourChest(yell string) {
	fmt.Println(yell + "!!!!")
}

func twoTypes(greet, name string) {
	fmt.Println(greet, name)
}

func sumVariatic(values ...int) int {
	result := 0

	for _, v := range values {
		result += v
	}
	return result
}

func sumVariaticPointer(values ...int) *int {
	result := 0

	for _, v := range values {
		result += v
	}
	return &result
}

func test(values int) (result int) {
	result = 1 + values
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

// basic syntax
/*
camelCase PascalCase
PascalCase is public
camelCase internal to package
func <name>(<params>) <return value> {

}

practice passing in pointers

funCall(&varName)

func funCall(*varName)

variatic params should be at the end of the func call and can only use one

values ...int <-- variatic function

go auto promotes funcs pointer from local stack to heap memory
most other languages it will die

*/
