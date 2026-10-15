// Struct Tags
// - Struct tags are used to attach metadata to struct fields.
// - They are commonly used to provide additional information to the encoding/json package.
// - Struct tags can be accessed by using reflection.
// - Syntax: `key:"value"`

package structures

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//==================================================================================================
// Declaring Tags
//==================================================================================================

// Declaring Struct with Tags
// - The struct below has a field with metadata ("tags") attached to it.

type Customer struct {
	Id int `json:"key"`
}

//==================================================================================================
// Using Tags
//==================================================================================================

func UsingStructTags() {
	customer := Customer{Id: 1}

	// Parsing Struct Tags (JSON)
	// - Tags are using by libraries as meta information for some operations.
	// - For example, in json library, the tag "json" is used to define the key name in the JSON
	//   output.
	// - Output: {"key":1}

	jsonData, _ := json.Marshal(customer)
	fmt.Println(string(jsonData))

	// Accessing Struct Tags
	// - Struct tags can be accessed using reflection.
	// - The "reflect" package is used to inspect the struct fields and their tags.
	// - Output: key

	val := reflect.TypeOf(customer).Field(0).Tag.Get("json")
	fmt.Println(val)
}
