package main

import "fmt"

type Salutation struct {
  name string
  greeting string
}

// Function Declaration:
// func <func_name>(arguments with types)
func Greet(salutation Salutation)  {
  fmt.Println(salutation.name)
  fmt.Println(salutation.greeting)
  // Calling Variadic function
  msg1, msg2 := GreetMultiple(salutation.name, salutation.greeting, "Hi")
  fmt.Println(msg1)
  fmt.Println(msg2)
}

// Calling a function inside another function
func GreetNew(salutation Salutation) {
  fmt.Println(CreateMessage(salutation.name, salutation.greeting))
  // Call createmessagemultiple [Similar to tuple unpacking in Python]
  message, alternate := CreateMessageMultiple(salutation.name, salutation.greeting)
  fmt.Println(message)
  fmt.Println(alternate)
  // In Go, Unused variables are errors and results in program termination.
  // In the above example, if only one greeting message is required from CreateMessageMultiple function, we can use '_' for dummy assignment
  _, alternate2 := CreateMessageMultiple(salutation.name, salutation.greeting)
  fmt.Println(alternate2)

}

// Function accepts two strings and returns a string
// func <func_name>(arguments [shorthand can be used]) <return type>
func CreateMessage(name, greeting string) string {
  return "greeting" + " " + name
}

// Multiple Returns
func CreateMessageMultiple(name, greeting string) (string, string) {
  return "greeting" + " " + name, "HEY!" + " " + name
  // return two messages
}

// Named Returns
// Return values can be named just like the arguments are defined resulting in a short function body
func CreateMessages(name, greeting string) (message string, alternate string)  {
  message = "greeting" + " " + name
  alternate = "HEY!" + " " + alternate
  return
  // Notice the above return keyword, all the variables declared to be returned are modified inside the function and the final return statement returns them all.
}

//Variadic Functions - functions which take indefinite number of arguments
// Use ...<type> in arguments
func GreetMultiple(name string, greeting ...string) (message, alternate string){
  fmt.Println(len(greeting)) // prints the length of the greeting string
  message = greeting[1] + " " + name
  alternate = greeting[0] + " " + name
  return

}

func main()  {
  var s = Salutation{"John", "Hello!"}
  Greet(s)
  GreetNew(s)
}
