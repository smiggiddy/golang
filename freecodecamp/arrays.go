package main

import (
	"fmt"
)

func main() {
	// declaring an array

	grades := [3]int{97, 85, 83} // number of elements must be declared and type of array needs to be stored
	fmt.Println(grades)

	// another way to declare an array
	grades_just_three := [...]int{1, 2, 3, 4, 5}
	fmt.Println(grades_just_three)

	// empty array
	var students [3]string
	fmt.Println(students)

	// there will be a pointer that shows where the array is in memory
	// head of array at position of 0 will move forward X elements of the string
	students[0] = "Bob"
	students[1] = "Bill"
	students[2] = "Drew"

	fmt.Println(students)
	fmt.Printf("Number of Students: %v\n", len(students))

	//array is considered a value, when you copy an array it creates a copy
	failures := students
	failures[1] = "Billy"

	fmt.Println(failures)
	fail := &students
	fmt.Println(fail)
	fail[1] = "Smiggiddy"
	fmt.Println(fail, students)

	//slice init as a literal

	slice := []int{1, 2, 3}
	fmt.Println(slice[1])
	failureZ := students
	failureZ[1] = "he really passed"
	fmt.Println(students)

	// using a make function type of object you'd like to create

	simple := make([]int, 3, 20)
	fmt.Println(simple)
	fmt.Printf("Length: %v\n", len(simple))
	fmt.Printf("Capacity: %v\n", cap(simple))

	simple = append(simple, 20)
	fmt.Println(simple)
	fmt.Printf("Length: %v\n", len(simple))
	fmt.Printf("Capacity: %v\n", cap(simple))

	// the 3rd element helps to expand the slice of the slice.. append it
	//variatic expression / function
	//shift op
	smaller_simple := simple[1:]
	fmt.Println(smaller_simple)

}
