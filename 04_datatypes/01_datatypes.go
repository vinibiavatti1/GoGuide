// Data Types
// - Go provides several built-in primitive data types.
// - The 'byte' type is an alias for uint8, and 'rune' is an alias for int32.
// - Numeric sizes and defaults can vary depending on the architecture (e.g., int).

package datatypes

// Boolean

var _ bool = true // true or false

// String

var _ string = "Hello" // UTF-8 string

// Integer

var _ int = 1     // 32/64-bit integer (default)
var _ int8 = 1    // 8-bit: -128 to 127
var _ int16 = 1   // 16-bit: -32768 to 32767
var _ int32 = 1   // 32-bit: -2147483648 to 2147483647
var _ int64 = 1   // 64-bit: -9223372036854775808 to 9223372036854775807
var _ uint = 1    // 32/64-bit Unsigned integer (default)
var _ uint8 = 1   // 8-bit: 0 to 255
var _ uint16 = 1  // 16-bit: 0 to 65535
var _ uint32 = 1  // 32-bit: 0 to 4294967295
var _ uint64 = 1  // 64-bit: 0 to 18446744073709551615
var _ uintptr = 1 // Memory Address Type (32 or 64 bits)
var _ byte = 1    // Alias for uint8
var _ rune = 'A'  // Alias for int32 (Unicode code point)

// Float

var _ float32 = 3.14 // 32-bit float
var _ float64 = 3.14 // 64-bit float (default)

// Complex

var _ complex64 = 1 + 2i  // 32-bit real and imaginary parts
var _ complex128 = 1 + 2i // 64-bit real and imaginary parts (default)

// Any

var _ any = 1 // Any stores values of any type
