// Type
// - The "type" keyword is used to declare new types or create aliases for existing types.
// - Syntax:
//   - New Type: type <name> <type>
//   - Alias:    type <name> = <type>

package typing

import "fmt"

//==================================================================================================
// Declaring Types
//==================================================================================================

// Declaring a Type
// - We can declare new types using the "type" keyword.
// - The new type can be based on an existing type.
// - In the example below, LocalType is a new type based on the int type.
// - Aliases can be created using the "=" operator.

type LocalType int   // New type
type AliasType = int // Alias (Using "=" operator)

// Declaring Multiple Types
// - We can use a type block to declare multiple types at once.
// - The examples below shows different base types that can be used for new types.

type (
	Number       float64            // New datatype
	Person       struct{}           // New Struct type
	Handler      interface{}        // New Interface type
	Sum          func(x, y int) int // New Func type
	Point        [2]int             // New Array type
	Collection   []any              // New Slice type
	Stack[T any] []T                // New Slice type (With Generics)
	Dictionary   map[any]any        // New Map type
	Decimal      = float64          // Alias (Using "=" operator)
)

//==================================================================================================
// Using Types
//==================================================================================================

func UsingTypes() {

	// Using Declared Struct
	// - Output: &{}

	person := &Person{}
	fmt.Println(person)

	// Using Declared Interface
	// - Output: &{}

	var handler Handler = &Person{}
	fmt.Println(handler)

	// Using Declared Function
	// - Output: 10

	var sum Sum = func(x, y int) int {
		return x + y
	}
	fmt.Println(sum(7, 3))

	// Using Declared Array
	// - Output: [1 2]

	p := Point{1, 2}
	fmt.Println(p)

	// Using Declared Slice
	// - Output: [A B C]

	c := Collection{"A", "B", "C"}
	fmt.Println(c)

	// Using Declared Map
	// - Output: map[A:1]

	d := Dictionary{"A": 1}
	fmt.Println(d)

	// Using New Type
	// - Output: 3.14

	var n Number = 3.14
	fmt.Println(n)

	// Using Alias
	// - Output: 3.14

	var x Decimal = 3.14
	fmt.Println(x)
}

//==================================================================================================
// Local Types
//==================================================================================================

func LocalTypes() {

	// Declaring Local Type
	// - We can declare a type within a function.
	// - This type is not accessible outside the function.

	type value int

	// Using Local Type
	// - We can use the local type within the function.
	// - Output: 42

	var v value = 42
	fmt.Println(v)
}
