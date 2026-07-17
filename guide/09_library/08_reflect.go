// Reflect
// The reflect package implements run-time reflection, allowing a program to
// inspect and manipulate objects with arbitrary types.

package library

import (
	"fmt"
	"reflect"
)

// Basic Reflection
// The following functions demonstrate how to inspect types and kinds of variables.
func BasicReflection() {
	// TypeOf vs Kind
	// TypeOf returns the exact static type (e.g., library.Person).
	// Kind returns the underlying specific variety of type (e.g., struct, int, string).
	// Output: Type: float64 | Kind: float64
	x := 3.14
	t := reflect.TypeOf(x)
	k := t.Kind()
	fmt.Printf("Type: %v | Kind: %v\n", t, k)

	// ValueOf
	// ValueOf returns a new Value initialized to the concrete value passed in.
	// Output: Value: 42
	v := reflect.ValueOf(42)
	fmt.Println("Value:", v)
}

// Struct Tags And Fields
// The following functions demonstrate how to inspect struct fields and read custom tags.
func StructTagsAndFields() {
	// Inspecting Structs and Tags
	// Reflection allows you to loop through struct fields and extract metadata like tags.
	// Output: Field: Name, Type: string, Tag: json:"user_name"
	// Output: Field: Age, Type: int, Tag: json:"user_age"
	type User struct {
		Name string `json:"user_name"`
		Age  int    `json:"user_age"`
	}

	u := User{Name: "Vini", Age: 30}
	t := reflect.TypeOf(u)

	// Loop through all fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		fmt.Printf("Field: %s, Type: %s, Tag: json:\"%s\"\n", field.Name, field.Type, tag)
	}
}

// Reflection Comparison
// The following functions show how reflection can be used to compare complex structures.
func ReflectionComparison() {
	// DeepEqual
	// DeepEqual reports whether x and y are "deeply equal".
	// Unlike the standard == operator, it can compare slices, maps, and deeply nested structs.
	// Output: true
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}

	// standard comparison slice1 == slice2 would fail to compile in Go
	isEqual := reflect.DeepEqual(slice1, slice2)
	fmt.Println(isEqual)
}
