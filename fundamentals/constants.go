package main
/*
Constants exists only at compile time
they use 256 bits of precision
*/

// constants are allowed to be based on builtin types
// Constant of a type
type Duration int16

const (
	Nanosecond Duration = 1
	Microsecond = 1000 * Nanosecond
)

func main()  {

	// These are constants of a kind
	// Untyped constants - can convert implicitly
	const pi = 3.1415 //kind: floating-point

	// Typed constants
	const abc int = 12345 // kind: int

	const (
		A = iota // starts with 0 inside const block
		B = iota // And then subsequent iota increments
		C = iota 
	)

	println("1: ", A, B, C)

	const (
		A1 = iota + 1 // starts 0 + 1 and subsequent can be skipped
		B1
		C1
	)

	println("2: ", A1, B1, C1)

}