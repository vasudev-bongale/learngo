package main

// Great place to understand and learn GO
// https://gobyexample.com/

import "fmt"

// https://blog.golang.org/go-slices-usage-and-internals

func main()  {

	// Declaration of slices
	s := make([]string, 3)
	fmt.Println("Slice: ", s)

	// Declare and Initialize slice in a single line
	t := []string{"good", "bad", "ugly"}
	fmt.Println("Slice with shorthand syntax: ", t)

	// Setting values
	s[0] = "x"
	s[1] = "y"
	s[2] = "z"

	// Appends values to an existing
	s = append(s, "w")
	fmt.Println("Slice with values set:", s)
	fmt.Println("Slice value at index 2: ", s[2])
	fmt.Println("Length of Slice: ", len(s))

	// Copying slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied slice C:", c)

	// Slices can be sliced similar to Python syntax
	fmt.Println("A slice of slice s", s[:3])

	// ... syntax expands an array
	s = append(s, s...) // s... : are the values of s
	fmt.Println("Slice s: ", s)
}