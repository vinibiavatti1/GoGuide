// Datatypes
// Go has several built-in data types, categorized as follows:
// - Boolean
// - String
// - Integer (Int, Byte and Rune)
// - Float
// - Complex

package datatypes

var (
	// Boolean
	_ bool = true // true or false

	// String
	_ string = "Hello" // UTF-8 string

	// Integer
	_ int     = 1   // 32/64-bit integer (default)
	_ int8    = 1   // 8-bit: -128 to 127
	_ int16   = 1   // 16-bit: -32768 to 32767
	_ int32   = 1   // 32-bit: -2147483648 to 2147483647
	_ int64   = 1   // 64-bit: -9223372036854775808 to 9223372036854775807
	_ uint    = 1   // 32/64-bit Unsigned integer (default)
	_ uint8   = 1   // 8-bit: 0 to 255
	_ uint16  = 1   // 16-bit: 0 to 65535
	_ uint32  = 1   // 32-bit: 0 to 4294967295
	_ uint64  = 1   // 64-bit: 0 to 18446744073709551615
	_ uintptr = 1   // Memory Address Type (32 or 64 bits)
	_ byte    = 1   // Alias for uint8
	_ rune    = 'A' // Alias for int32 (Unicode code point)

	// Float
	_ float32 = 3.14 // 32-bit float
	_ float64 = 3.14 // 64-bit float (default)

	// Complex
	_ complex64  = 1 + 2i // 32-bit real and imaginary parts
	_ complex128 = 1 + 2i // 64-bit real and imaginary parts (default)

	// Any
	_ any = 1 // Any stores values of any type
)
