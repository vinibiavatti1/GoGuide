// Structs
// - Structs are composite data types that group together variables (fields) under a single name.
// - They are used to create complex data types that can represent real-world entities.
// - Structs can contain fields of different types, including other structs, arrays, and slices.

package structures

import "fmt"

//==================================================================================================
// Declaring Structs
//==================================================================================================

// Defining a Struct
// - We use the "struct" keyword to define structs.
// - The struct below has two fields: "Name" and "surname".
// - The first letter of the field name determines its visibility: if it is capitalized, it is
//   public, if not, it is private.
type Person struct {
	Name    string // Public
	surname string // Private (Not capitalized)
}

// Composition
// - Composition is a type of relationship where one struct can include another as a field.
// - Here, the "Employee" struct includes a pointer to the "Person" struct.
// - This allows the "Employee" to access the "Person" fields, along with its own additional field,
//   like the salary.
// - We can omit the name of the composition to use the same name of the type.
type Employee struct {
	*Person // Pointer to Person, allowing direct access to Name and Surname fields
	Salary  float64
}

// Constructors
// - Constructors are methods responsible for creating struct instances.
// - As a convention, they are named "New" followed by the struct name.
// - They are used to initialize the struct fields and return a pointer to the struct.
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
// - Methods are functions that are associated with a struct type.
// - They are defined using the "func" keyword, followed by the receiver type (the struct type) and
//   the method name.
// - In this example, the method "FullName" is associated with the "Employee" struct and takes a
//   pointer receiver (e *Employee).
func (e *Employee) FullName() string {
	return e.Name + e.surname
}

// Getterns and Setters
// - Getters and setters are methods that allow access to private fields.
// - In this example, the "Surname" method is a getter for the "surname" field in the "Person"
//   struct.
// - The "SetSurname" method is a setter for the "surname" field.
// - Note: In Go, we don't use the "Get" prefix for getters.
func (p *Person) Surname() string { // Getter
	return p.surname
}
func (p *Person) SetSurname(surname string) { // Setter
	p.surname = surname
}

//==================================================================================================
// Using Structs
//==================================================================================================

func UsingStructs() {

	// Creating an Instance
	// - Creating an instance of "Employee", initializing the "Person" field using a pointer.
	employee1 := &Employee{
		Person: &Person{
			Name:    "John",
			surname: "Duo",
		},
		Salary: 2000.0,
	}

	// Creating an Instance Using Constructor Method
	// - We can also use the constructor method to create an instance of Employee.
	employee2 := NewEmployee("John", "Duo", 2000.0)

	// Accessing Public Data
	// - Embedded fields can be accessed directly from the "Employee" struct.
	// - Note that the "Name" field belongs to the "Person" struct, but it can be accessed directly
	//   from the "Employee" struct.
	name1 := employee1.Person.Name // Accessing emdebbed
	name2 := employee1.Name        // Accessing directly
	fmt.Println(name1, name2)      // Output: John John

	// Calling Methods
	// - Calling the "FullName" method associated to the employee.
	// - Output: fullname: JohnDuo
	fullname := employee2.FullName()
	fmt.Println("fullname:", fullname)

	// Acessing Private Data
	// - Since the "Surname" field is private, it can be accessed only by public methods.
	// - Note that the "Surname" method belongs to the "Person" struct, but it can be accessed
	//   directly from the "Employee" struct.
	// - Output: Duo Duo
	surname1 := employee1.Person.Surname() // Accessing emdebbed
	surname2 := employee1.Surname()        // Accessing directly
	fmt.Println(surname1, surname2)

	// Mutating Data
	// - We can assign a new value to a public field of the struct.
	// - We can also set the value of embedded fields directly.
	// - Output: Salary: 2500.0
	employee1.Salary = 2500.0   // Setting Salary
	employee1.Name = "John Doe" // Setting Name directly
	fmt.Println("Salary:", employee1.Salary)

	// Local Structs
	// - We can declare a struct inside a function.
	// - This struct will only be available inside the function.
	type user struct {
		Name string
	}

	// Creating an Instance of Local Struct
	// - The local struct is only available inside the function
	// - Output: otherPerson: {John Duo}
	_user := &user{
		Name: "John Duo",
	}
	fmt.Println("otherPerson:", _user)

	// Anonymous Structs
	// - We can create anonymous structs, which are structs without a name.
	// - They are useful for quick and temporary data structures.
	// - Note that the struct below is not named, but it has fields.
	// - Output: person: {John Duo}
	person := struct {
		Name string
	}{
		Name: "John Duo",
	}
	fmt.Println("person:", person)
}
