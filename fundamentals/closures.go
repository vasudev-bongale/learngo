package main

import "fmt"

type Salutation struct {
  name string
  greeting string
}

// Functions are types in Go
// Declaring a type as a function with a single string argument
type Printer func(string) ()

// Passing another function as argument to a function
// Can be passed as "do func(string)" or
// Type can be defined outside and passed
func Greet(salutation Salutation, do Printer) {
  message, alternate := CreateMessage(salutation.name, salutation.greeting)
  do(message)
  do(alternate)
}

func CreateMessage(name, greeting string) (message string, alternate string) {
  message = greeting + " " + name
  alternate = "HEY!" + " " + name
  return
}

func Print(s string) {
  fmt.Print(s)
}

func PrintLine(s string) {
  fmt.Println(s)
}

/*
Closure - returns a modified function by imprinting the variables from the context
in which the function is defined into the function which is being created.
Example: if we want to pass a custom string along with the string to different functions.
*/

// Returns a Printer Function
func CreatePrinter(custom string) Printer {
  return func(s string)  {
    fmt.Println(s + custom)
  }
}

func main()  {
  var s = Salutation{"John", "Hello!"}
  // A function parameter is passed, which alters the behaviour of the original Greet function
  Greet(s, PrintLine)
  Greet(s, Print)
  // CreatePrinter takes a custom argument and returns a function of type Printer by adding the custom string
  Greet(s, CreatePrinter("!!!"))

}
