package main

import "fmt"

/*
Method is a function that has a receiver. Value receiver or pointer receiver.
When do we use value vs pointer recievers in methods?

General guidelines
There are 3 different categories of types in go:
1. built-in types - (Value semantics)  includes fields in structs
2. Reference types - slices, maps, etc (Value semantics)
3. User defined types - structs () (pointer semantics)

Decide whether your data need behavior (in the form of methods)

functions are values in Go.
*/

type user struct {
	name  string
	email string
}

func main() {
	bill := user{"bill", "bill@nowhere.com"}

	// Although changeEmail uses pointer reciever,
	// go automatically uses the pointer for bill.
	bill.changeEmail("newemail@nowhere.com")
	bill.notify()

	// Functions are values in go
	// This is decoupling. It allows a level of indirection.
	// The cost of this is allocation, as the escape analysis
	// is not sure anymore that the data should be on stack.
	f1 := bill.notify
	f1()

	// Now Lets change the email of bill
	bill.changeEmail("updated@nowhere.com")

	// This will still print the original email of bill as of line 24
	// This is because notify uses value semantics and has its own
	// copy of the data
	f1()

	f2 := bill.changeEmail
	bill.notify()

	// Lets change email again
	f2("email@now.com")
	bill.notify()

}

// A method on user with Value reciever
func (u user) notify() {
	fmt.Printf("Sending an email to %s<%s>\n", u.name, u.email)
}

// A method on user with pointer receiver
func (u *user) changeEmail(email string) {
	u.email = email
}
