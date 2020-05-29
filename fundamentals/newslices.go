package main

/*
Most important data structure in Go.

make - built-in function in go that can be used to create referenced types.
such as slices, maps

A slice has a length and capacity.

Length = total number elements that can be reached from the pointer position
capacity = total number of elements that exists from that pointer in that backing element.

capacity can be larger than lenght but not other way.
*/

func main()  {

	// This is a nil slice
	var data []string

	// This literal construction returns an empty slice
	// data := []string{}
	// If this is an empty slice, where is it pointing to?
	// Go has an inbuilt structure which is a zero allocation struct.
	// var es struct{}
	
	// Create a slice with a length of 5 elements.
	fruits := make([]string, 5)

	// Create a slice with len5, cap8
	vegetables := make([]string, 5, 8)

}