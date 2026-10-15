// Error
// - In Go, there isn't a try-catch mechanism, instead, errors are handled by checking the error
//   value returned from a function.
// - Errors in Go are values that implement the error interface.
// - The error value is nil if there is no error.

package errors

import (
	"errors"
	"fmt"
)

//==================================================================================================
// Handling Errors
//==================================================================================================

// Declaring Function that Raises an Error
// - The function below shows how to declare a function that returns an error.
// - If we need a formatted error, we can use the "fmt.Errorf" function.
// - The error can be nil if there is no error.
// - It returns an error value that can be used to format the error message.

func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero %d/%d", x, y)
	}
	return x / y, nil
}

// Handling Errors
// - We can check if the error value is nil or not.
// - If the error value is not nil, we can handle the error.
// - Output: Error: division by zero 4/0

func HandlingError() {
	r, err := Divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", r)
	}
}

//==================================================================================================
// Wrapping Errors
//==================================================================================================

// Wrapping an Error
// - Wrapping an error means to append the current error to the new error.
// - This is useful to add more context to the error.
// - We can wrap an error using the "fmt.Errorf" function with the "%w" (wrap) verb.
// - The function below performs the other function, and wraps the error if it occurs.

func PerformDivision(x, y int) (int, error) {
	r, err := Divide(x, y)
	if err != nil {
		return 0, fmt.Errorf("failed to divide: %w", err)
	}
	return r, nil
}

// Handling Wrapped Error
// - Now, when the errors is printed, it will show the wrapped error and the original error.
// - This way is more recommended since it is easier to locate the error.
// - Output: Error: failed to divide: division by zero 4/0

func HandlingWrappedError() {
	_, err := PerformDivision(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successful")
	}
}

//==================================================================================================
// Custom Errors
//==================================================================================================

// Creating a Custom Error
// - Any struct that implements the error interface can be used as a custom error.
// - The "error" interface is located in the "errors" package.

type CustomError struct {
	Code    int
	Message string
}

// Implementing the Error Interface
// - The "Error()" method is used to implement the error interface.

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Declaring Function that can Raises a Custom Error
// - The function below raises the custom error.

func Request(path string) (string, error) {
	if path == "/admin" {
		return "", &CustomError{
			Code:    403,
			Message: "Forbidden",
		}
	}
	return "Success", nil
}

// Handling a Custom Error
// - The function below shows how to handle the custom error.
// - We can check the type of the error using type assertion to ensure that the error is of type
//   "CustomError".
// - Output: Error: Code: 403, Message: Forbidden

func HandlingCustomError() {
	_, err := Request("/admin")
	if err != nil {
		if _, ok := err.(*CustomError); ok {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Generic Error")
		}
	} else {
		fmt.Println("Successful")
	}
}

//==================================================================================================
// Sentinel Errors
//==================================================================================================

// Defining Sentinel Errors
// - Sentinel errors are static error values defined at the package level.
// - By convention, their names start with "Err". They are used when the caller only needs to know
//   if a specific error happened, without dynamic context.
// - For static strings, "errors.New" is preferred over "fmt.Errorf".

var ErrNotFound = errors.New("resource not found")
var ErrPermissionDenied = errors.New("permission denied")

// Declaring Function that Returns a Sentinel Error
// - The function below demonstrates how to return a predefined sentinel error.
// - This is useful when the error condition is predictable and static.

func FetchData(id int) (string, error) {
	if id < 0 {
		return "", ErrNotFound
	}
	return "data", nil
}

//==================================================================================================
// Error Utilities
//==================================================================================================

// Handling Errors with errors.Is
// - We should use "errors.Is" instead of the "==" operator to check for sentinel errors.
// - The "errors.Is" is smarter because it can inspect wrapped errors (created with %w).
// - The code below wraps the sentinel error to simulate a real-world multi-layer architecture.
// - Even though the error is wrapped, errors.Is can find ErrNotFound inside it.
// - Output: Error: Caught a NotFound error!

func HandlingWithErrorsIs() {
	_, err := FetchData(-1)
	wrappedErr := fmt.Errorf("database query failed: %w", err)
	if errors.Is(wrappedErr, ErrNotFound) {
		fmt.Println("Error: Caught a NotFound error!")
	}
}

// Handling Custom Errors with errors.As
// - In your previous example, you used type assertion (err.(*CustomError)).
// - However, if a CustomError is wrapped using "%w", type assertion will FAIL.
// - The "errors.As" is the recommended way because it unwraps the error chain to find the target
//   type.
// - Output: Error: Found CustomError with Code 403

func HandlingWithErrorsAs() {
	_, err := Request("/admin")
	wrappedErr := fmt.Errorf("api gateway error: %w", err)
	var myCustomErr *CustomError
	if errors.As(wrappedErr, &myCustomErr) {
		fmt.Printf("Error: Found CustomError with Code %d\n", myCustomErr.Code)
	}
}
