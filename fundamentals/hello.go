package main

import "fmt"

// Struct - collection of types or fields used to create user defined types
// There are no classes in go.
// Notice captial 'S' in the name Salute - indicates that the type can be accessible outside the module when imported, which is public.
type Salute struct {
  name string
  message string
}

// Defining constants
const (
  PI = 3.14
  Language = "Go"
)

// iota: represents a successive incremental unsigned integer
// Notice that B and C are not assigned any value, which by Go's behaviour takes the previous declaration, which is iota itself.
const (
  A = iota
  B
  C
)

func main() {

  //Declaration and initialization
  var var_name string = "Nice"
  //var var_name2 = "Auto" // Types can be automatically identified by Go
  //var_name3 := "Good" // Another form which can be used in the main function

  // Pointers: Holds the memory address to the variable Type
  // Pointer of a particular type cannot be used for another type
  var mypointer *string = &var_name // mypointer is a pointer type pointing to the address '&' of variable var_name2
  *mypointer = "Great" // *pointer_name accesses the value at the pointed address
  fmt.Println(mypointer) // prints the memory address of var_name
  fmt.Println(*mypointer) // prints value 'Great'

  //Declaration of user defined types
  // var message = Salute{} //valid declaration
  var message = Salute{message: "Joe", name: "Hello!"} // keyword declaration
  message.name = *mypointer // from the Salute Struct, the fields can be individually assigned
  message.message = "Hello!"

  fmt.Println(message)
  fmt.Println(Language, A, B, C)
}
