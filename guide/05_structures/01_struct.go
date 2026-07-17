// Struct
// Structs are composite data types that group together variables (fields) under a single name.
// They are used to create complex data types that can represent real-world entities.
// Structs can contain fields of different types, including other structs, arrays, and slices.

package datatypes

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Defining a Struct
// We use the "struct" keyword to define structs.
// The struct below has two fields: "Name" and "surname".
// The first letter of the field name determines its visibility: if it is capitalized, it is public; if not, it is private.
type Person struct {
	Name    string // Public
	surname string // Private (Not capitalized)
}

// Composition
// Composition is a type of relationship where one struct can include another as a field.
// Here, the "Employee" struct includes a pointer to the "Person" struct.
// This allows the "Employee" to access the "Person" fields, along with its own additional field, like the salary.
// We can omit the name of the composition to use the same name of the type.
type Employee struct {
	*Person // Pointer to Person, allowing direct access to Name and Surname fields
	Salary  float64
}

// Constructors
// Constructors are methods responsible for creating struct instances.
// As a convention, they are named "New" followed by the struct name.
// They are used to initialize the struct fields and return a pointer to the struct.
func NewEmployee(name, surname string, salary float64) *Employee {
	return &Employee{
		Person: &Person{
			Name:    name,
			surname: surname,
		},
		Salary: salary,
	}
}

// Methods
// Methods are functions that are associated with a struct type.
// They are defined using the "func" keyword, followed by the receiver type (the struct type) and the method name.
// In this example, the method "FullName" is associated with the "Employee" struct and takes a pointer receiver (e *Employee).
func (e *Employee) FullName() string {
	return e.Name + e.surname
}

// Getterns and Setters
// Getters and setters are methods that allow access to private fields.
// In this example, the "Surname" method is a getter for the "surname" field in the "Person" struct.
// The "SetSurname" method is a setter for the "surname" field.
// Note: In Go, we don't use the "Get" prefix for getters.
func (p *Person) Surname() string { // Getter
	return p.surname
}
func (p *Person) SetSurname(surname string) { // Setter
	p.surname = surname
}

// Using Structs
// This function demonstrates how to create an instance of a struct, access its fields, mutate data, and call methods.
func UsingStructs() {

	// Creating an Instance
	// Creating an instance of "Employee", initializing the "Person" field using a pointer.
	employee1 := &Employee{
		Person: &Person{
			Name:    "John",
			surname: "Duo",
		},
		Salary: 2000.0,
	}

	// Creating an Instance Using Constructor Method
	// We can also use the constructor method to create an instance of Employee.
	employee2 := NewEmployee("John", "Duo", 2000.0)

	// Accessing Public Data
	// Embedded fields can be accessed directly from the "Employee" struct.
	// Note that the "Name" field belongs to the "Person" struct, but it can be accessed directly from the "Employee" struct.
	name1 := employee1.Person.Name // Accessing emdebbed
	name2 := employee1.Name        // Accessing directly
	fmt.Println(name1, name2)      // Output: John John

	// Calling Methods
	// Calling the "FullName" method associated to the employee.
	fullname := employee2.FullName()
	fmt.Println("fullname:", fullname) // Output: fullname: JohnDuo

	// Acessing Private Data
	// Since the "Surname" field is private, it can be accessed only by public methods.
	// Note that the "Surname" method belongs to the "Person" struct, but it can be accessed directly from the "Employee" struct.
	surname1 := employee1.Person.Surname() // Accessing emdebbed
	surname2 := employee1.Surname()        // Accessing directly
	fmt.Println(surname1, surname2)        // Output: Duo Duo

	// Mutating Data
	// We can assign a new value to a public field of the struct.
	// We can also set the value of embedded fields directly.
	employee1.Salary = 2500.0                // Setting Salary
	employee1.Name = "John Doe"              // Setting Name directly
	fmt.Println("Salary:", employee1.Salary) // Output: Salary: 2500.0

	// Local Structs
	// We can declare a struct inside a function.
	// This struct will only be available inside the function.
	type user struct {
		Name string
	}

	// Creating an Instance of Local Struct
	// The local struct is only available inside the function
	_user := &user{
		Name: "John Duo",
	}
	fmt.Println("otherPerson:", _user) // Output: otherPerson: {John Duo}

	// Anonymous Structs
	// We can create anonymous structs, which are structs without a name.
	// They are useful for quick and temporary data structures.
	// Note that the struct below is not named, but it has fields.
	person := struct {
		Name string
	}{
		Name: "John Duo",
	}
	fmt.Println("person:", person) // Output: person: {John Duo}2
}

// Struct Tags
// Struct tags are used to attach metadata to struct fields.
// They are commonly used to provide additional information to the encoding/json package.
// Syntax: `key:"value"`
type Customer struct {
	Id int `json:"key"`
}

// Using Struct Tags
// This function demonstrates how to use struct tags to provide metadata to fields.
func UsingStructTags() {

	// Creating an Instance
	customer := Customer{
		Id: 1,
	}

	// Parsing Struct Tags (JSON)
	// Tags are using by libraries as meta information for some operations.
	// For example, in json library, the tag "json" is used to define the key name in the JSON output.
	jsonData, _ := json.Marshal(customer)
	fmt.Println(string(jsonData)) // Output: {"key":1}

	// Accessing Struct Tags
	// Struct tags can be accessed using reflection.
	// The "reflect" package is used to inspect the struct fields and their tags.
	val := reflect.TypeOf(customer).Field(0).Tag.Get("json")
	fmt.Println(val) // Output: key
}
