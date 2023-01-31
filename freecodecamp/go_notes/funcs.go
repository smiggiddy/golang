package main

import (
	"fmt"
)

func deezNuts(a, b int) int {
	return a + b
}

func madNumbers(nums ...int) int {
	total := 0

	for _, nums := range nums {
		total += nums
	}
	fmt.Println("Total nums:", total)
	return total
}

// recursion example
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	sum := deezNuts(3, 4)
	fmt.Printf("Sum: %v\n", sum)

	madNumbers(1, 2, 3, 4)

	fmt.Println(fact(4))
}
