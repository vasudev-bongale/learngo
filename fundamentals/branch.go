package main

import "fmt"

func GreetPeople(name string, isFormal bool)  {
  message := CreateMessage(name, isFormal)
  fmt.Println(message)
}

func CreateMessage(name string, isFormal bool) (message string) {
  // Basic if construct.
  // Embedded statements - can be used for limited scope execution
  // example: 'Mr' Prefix can be defined for the if loop in this way
  surname := GetSurname(name)
  if prefix := GetPrefix(name); isFormal {
    message = "Hello! " + name + " " + surname
  } else {
    message = "Hello! " + prefix + name + " " + surname
  }
  return
}

// Switch construct
func GetPrefix(name string) (prefix string) {
  switch name {
    case "Bob":
      prefix = "Mr "
      fallthrough // Fallthrough allows the case to match the next statment if it matches
      // ex: if Bob is the name, then prefix shall be "Dr" instead of Mr.
    case "John", "Peter": // Multiple values can be matches at once
      prefix = "Dr "
    case "Mary": prefix = "Mrs "
    default: prefix = "Dude "
  }
  return
}
//Switches can switch based on expression as well
func GetSurname(name string) (surname string) {
  // Switch on Nothing
  switch {
    case name == "John":
      surname = "Williams"
    case name == "Mary":
      surname = "Baggins"
    default:
      surname = "Unknowns"
  }
  return
}

// Switch on type
func TypeSwitch(x interface{})  {
  switch x.(type) {
    case int:
      fmt.Println("int")
    case string:
      fmt.Println("string")
    default:
      fmt.Println("Unknown")
  }
}

func main()  {
    GreetPeople("John", false)
    GreetPeople("Vasu", false)
    TypeSwitch(1)
    TypeSwitch(1.3)
}
