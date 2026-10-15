// Reflect
// - The reflect package implements run-time reflection, allowing a program to inspect and
//   manipulate objects with arbitrary types at runtime.

package library

import (
	"fmt"
	"reflect"
)

//==================================================================================================
// Core Functions
//==================================================================================================

func ReflectionCoreFunctions() {

	// TypeOf
	// - TypeOf returns the reflect.Type that represents the dynamic type of the interface.
	// - Output: main.User

	type User struct{ Name string }
	u := User{Name: "John"}
	t := reflect.TypeOf(u)
	fmt.Println(t)

	// ValueOf
	// - ValueOf returns a new Value initialized to the concrete value stored in the interface.
	// - Output: {John}

	v := reflect.ValueOf(u)
	fmt.Println(v)
}

//==================================================================================================
// Type Functions
//==================================================================================================

func ReflectTypeFunctions() {

	// Kind
	// - Kind returns the specific underlying variety of a type (e.g., struct, int, slice).
	// - Output: struct

	type User struct{ Name string }
	t := reflect.TypeOf(User{})
	fmt.Println(t.Kind())

	// NumField
	// - NumField returns a struct type's field count. It panics if the type's Kind is not Struct.
	// - Output: 2

	type Order struct {
		ID    int
		Total float64
	}
	t = reflect.TypeOf(Order{})
	fmt.Println(t.NumField())

	// Field
	// - Field returns the i'th field of a struct type. It panics if the index is out of bounds.
	// - Output: ID int

	f := t.Field(0)
	fmt.Println(f.Name, f.Type)

	// FieldByName
	// - FieldByName returns the struct field with the given name and a boolean indicating if it was
	//   found.
	// - Output: Total true

	f, found := t.FieldByName("Total")
	fmt.Println(f.Name, found)
}

//==================================================================================================
// Tag Functions
//==================================================================================================

func ReflectTagFunctions() {

	// Get
	// - Get returns the value associated with key in the tag string.
	// - Output: user_id

	type Profile struct {
		ID string `json:"user_id" db:"id"`
	}
	t := reflect.TypeOf(Profile{})
	field := t.Field(0)
	fmt.Println(field.Tag.Get("json"))

	// Lookup
	// - Lookup returns the value associated with key in the tag string and whether it was found.
	// - Output: id true

	val, ok := field.Tag.Lookup("db")
	fmt.Println(val, ok)
}

//==================================================================================================
// Value Functions
//==================================================================================================

func ReflectValueMethods() {

	// Interface
	// - Interface returns v's current value as an interface{}, packing it back into standard Go
	//   data.
	// - Output: 100

	v := reflect.ValueOf(100)
	i := v.Interface().(int) // Type assertion back to int
	fmt.Println(i)

	// Elem
	// - Elem returns the value that the interface v contains or that the pointer v points to.
	// - Essential when you need to modify a value via a pointer.
	// - Output: 200

	x := 100
	v = reflect.ValueOf(&x) // Passing pointer
	v.Elem().SetInt(200)    // Modifying the underlying value
	fmt.Println(x)

	// IsNil
	// - IsNil reports whether its argument v is nil. v must be a chan, func, interface, map,
	//   pointer, or slice.
	// - Output: true

	var s []string
	v = reflect.ValueOf(s)
	fmt.Println(v.IsNil())

	// IsValid
	// - IsValid reports whether v represents a value. It returns false if v is the zero Value (like
	//   from a missing map key).
	// - Output: false

	v = reflect.ValueOf(nil)
	fmt.Println(v.IsValid())
}

//==================================================================================================
// Utility Functions
//==================================================================================================

func ReflectionUtilityFunctions() {

	// DeepEqual
	// - DeepEqual reports whether x and y are "deeply equal", working natively on slices, maps, and
	//   nested structs.
	// - Output: true

	m1 := map[string]int{"A": 1}
	m2 := map[string]int{"A": 1}
	fmt.Println(reflect.DeepEqual(m1, m2))

	// Copy
	// - Copy copies elements from a source slice to a destination slice using reflection Values.
	// - Output: 2 (elements copied) -> [1 2]

	src := reflect.ValueOf([]int{1, 2, 3})
	dst := reflect.ValueOf(make([]int, 2))
	copied := reflect.Copy(dst, src)
	fmt.Println(copied, dst.Interface())
}
