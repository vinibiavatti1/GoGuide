// Recursion
// Go supports recursion, which is a function calling itself.
// Recursion is a powerful technique for solving problems that can be broken down into smaller subproblems.

package gopatterns

import "fmt"

// Recursive Factorial
// We can use recursion to calculate the factorial of a number.
// The factorial of a number n is the product of all positive integers less than or equal to n.
func recursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

// Perform Recursive Factorial
// This function demonstrates the use of recursion to calculate the factorial of a number.
func PerformRecursiveFactorial() {

	// Perform Recursive Factorial
	// We will call the recursive function to show the result.
	fmt.Println("Factorial of 5 is:", recursiveFactorial(5)) // Output: Factorial of 5 is: 120
	fmt.Println("Factorial of 0 is:", recursiveFactorial(0)) // Output: Factorial of 5 is: 1
}
