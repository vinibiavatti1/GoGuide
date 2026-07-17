// Goto
// The "goto" statement allows jumping to a labeled statement within the same function.
// It is generally discouraged in favor of structured control flow (loops, conditionals, etc.).
// However, it can be useful for breaking out of deeply nested structures.
// Syntax:
// - goto <label>
// - <label>:

package syntax

import "fmt"

// Simple Goto Example
// Demonstrates a basic jump using goto.
// Output: Start | End
func SimpleGoto() {
	fmt.Println("Start")
	goto skip
	// Unreached code
skip:
	fmt.Println("End")
}

// Loop Simulation with Goto
// Shows how goto can be used for a manual loop-like structure.
// The "goto" statement below jumps back to "repeat" label, acting like a loop.
// Output: 0 | 1 | 2
func LoopWithGoto() {
	x := 0
repeat:
	fmt.Println(x)
	x++
	if x < 3 {
		goto repeat
	}
}

// Using Goto for Error Handling
// Goto can be used to centralize error handling.
// Output (if err = true): An error occurred. Handling error...
func ErrorHandlingWithGoto() {
	err := false
	if err {
		goto err
	}
	fmt.Println("Operation successful")
	return
err:
	fmt.Println("An error occurred. Handling error...")
}

// Breaking Out of Loops
// Goto can be used to break out of loops.
// It is useful for breaking out of deeply nested loops for example.
// Output: 0 | 1 | 2 | Exited nested loops
func BreakingOutLoops() {
	for i := range 10 {
		if i == 3 {
			goto exit
		}
		fmt.Println(i)
	}
exit:
	fmt.Println("Exited nested loops")
}
