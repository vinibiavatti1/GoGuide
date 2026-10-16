// Random
// - The math/rand package implements pseudo-random number generators.
// - It is useful for tasks such as simulations, testing, and generating random data.
// - Note: For security-sensitive applications, use the "crypto/rand" package instead.

package library

import (
	"fmt"
	"math/rand"
)

//==================================================================================================
// Basic Random Functions
//==================================================================================================

func RandomBasicFunctions() {

	// IntN
	// - IntN returns, as an int, a non-negative pseudo-random number in the half-open interval
	//   [0, n).
	// - Panics if n <= 0.
	// - Output: n: 4 (Value between 0 and 9)

	n := rand.Intn(10)
	fmt.Println("n:", n)

	// Float64
	// - Float64 returns, as a float64, a pseudo-random number in the half-open interval
	//   [0.0, 1.0).
	// - Output: f: 0.7492...

	f := rand.Float64()
	fmt.Println("f:", f)

	// Perm
	// - Perm returns, as a slice of n integers, a pseudo-random permutation of the integers in
	//   [0, n).
	// - Output: perm: [2 0 1]

	perm := rand.Perm(3)
	fmt.Println("perm:", perm)

	// Shuffle
	// - Shuffle pseudorandomizes the order of elements using the provided swap function.
	// - Note: It modifies the original slice in place.
	// - Output: x: [C A B]

	x := []string{"A", "B", "C"}
	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})
	fmt.Println("x:", x)
}

//==================================================================================================
// Custom Source Functions
//==================================================================================================

func RandomCustomSource() {

	// NewRand with Custom Source
	// - We can create a custom local pseudo-random number generator using a seeded source.
	// - This avoids sharing global generator state and allows reproducible results.
	// - Output: val: 42

	src := rand.NewSource(99)
	rng := rand.New(src)
	val := rng.Intn(100)
	fmt.Println("val:", val)
}
