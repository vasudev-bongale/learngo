package main

import "fmt"
/*
Arrays are a collections of objects of same type
Arrays:
* Are fixed length
* Array type is the size and the underlying type [ array with size 3 and of int is of different type than array
  with size 4 and type 4
* Arrays are passed by Values
* No intialization ( 0 valued )


Slices:
* Are not fixed sized (can be expanded by appending)
* Slice type is the underlying type
* Slice is a pointer
* Initialized null values (make)
 */

func main() {
	// Intializing arrays
	var s [3]int
	// Updating the array
	s[2] = 100
	fmt.Println("Array:", s)
	fmt.Println("Length of the Array:" , len(s))

	c := [4]int{1, 2, 3, 4}
	fmt.Println("Array C:", c)

	// 2D arrays
	var twoDarray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDarray[i][j] = i * j
		}
	}
	fmt.Println("TwoDArray: ", twoDarray)
}