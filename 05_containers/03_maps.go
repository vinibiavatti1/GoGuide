// Maps
// - Maps are a collection of key-value pairs.
// - Maps are unordered collections, meaning the order of elements is not guaranteed.
// - Syntax: map[<keyType>]<valueType>{<key>:<value>}

package containers

import "fmt"

//==================================================================================================
// Declaring Maps
//==================================================================================================

func DeclaringMaps() {

	// Declaring a Map
	// - Maps are declared using the "map" keyword followed by the key and value types.
	// - The key type must be a comparable type, while the value type can be any type.
	// - When no elements are specified, the map will be empty.
	// - Output: x: map[]

	x := map[string]int{}
	fmt.Println("x:", x)

	// Declaring a Map with Predefined Values
	// - We can initialize the map with predefined key-value pairs.
	// - The order of elements is not guaranteed.
	// - Output: x: map[A:1 B:2]

	x = map[string]int{
		"A": 1,
		"B": 2,
	}
	fmt.Println("x:", x)

	// Declaring a Map of "any" Type
	// - We can declare a map where the key and value types are of type "any".
	// - The "any" type is equivalent to the "interface{}" type in Go.
	// - Output: y: map[A:1 3.4:true]

	y := map[any]any{
		"A": 1,
		3.4: true,
	}
	fmt.Println("y:", y)
}

//==================================================================================================
// Map Operations
//==================================================================================================

func MapOperations() {
	x := map[string]int{"A": 1, "B": 2, "C": 3}

	// Acessing Elements by Key
	// - We access elements in a map using the key.
	// - If the key is not present, the zero value of the value type is returned.
	// - Output: y: 1 z: 0

	y := x["A"]
	z := x["?"]
	fmt.Println("y:", y, "z:", z)

	// Validating Key Existence
	// - When accessing elements, we can check if the key exists in the map by looking at the second
	//   return value.
	// - If the key is not present, the second return value will be false.
	// - Output: v: 0 ok: false

	v, ok := x["?"]
	fmt.Println("v:", v, "ok:", ok)

	// Retrieving Length
	// - We can get the number of key-value pairs in a map using the "len()" function.
	// - Output: Len: 3

	l := len(x)
	fmt.Println("Len:", l)

	// Adding Elements
	// - We can add key-value pairs to a map by assigning a value to a new key.
	// - Output: x: map[A:1 B:2 C:3 D:9]

	x["D"] = 9
	fmt.Println("x:", x)

	// Updating Elements
	// - We can update the value of an existing key by assigning a new value to it.
	// - Output: x: map[A:1 B:2 C:3 D:4]

	x["D"] = 4
	fmt.Println("x:", x)

	// Deleting Elements
	// - We can use the "delete()" builtin function to remove a key-value pair from a map.
	// - Output: x: map[A:1 B:2 C:3]

	delete(x, "D")
	fmt.Println("x:", x)

	// Clearing a Map
	// - We can use the "clear()" builtin function to remove all key-value pairs from a map.
	// - Output: x: map[]

	clear(x)
	fmt.Println("x:", x)

	// Iterating over Maps Entries
	// - We can use the for-range loop to iterate over the key-value pairs in a map.
	// - Output: A 1 | B 2 | C 3

	for k, v := range x {
		fmt.Println(k, v)
	}

	// Iterating over Maps Keys
	// - To iterate over the keys of a map, we can define only the key in the for-range loop.
	// - Output: A | B | C

	for k := range x {
		fmt.Println(k)
	}

	// Iterating over Maps Values
	// - To iterate over the values of a map, we can use the blank identifier "_" for the key.
	// - Output: 1 | 2 | 3

	for _, v := range x {
		fmt.Println(v)
	}
}
