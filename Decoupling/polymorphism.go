package main

import "fmt"

/*
"Polymorphism means that you write a certain program and it behaves differently
depending on the data that it operates on.""
More specifically, depending on concrete data.

Method sets:
- For a type T (non pointer) only funcs with value receivers can form its methods set
- For a type *T both pointer and value receivers
*/

// Interfaces are not real/concrete. They only define a method set of behaviour (verb)
type reader interface {
	read(b []byte) (int, error)
}

type file struct {
	name string
}

type pipe struct {
	name string
}

// read implements the reader interface for a file
func (file) read(b []byte) (int, error) {
	s := "<html>Hello Go</html>"
	copy(b, s)
	return len(s), nil
}

// read implements the reader interface for a file
func (pipe) read(b []byte) (int, error) {
	s := "<html>Hello Go From Pipe</html>"
	copy(b, s)
	return len(s), nil
}

// Interface types are valueless.
// Here it says, pass any concrete data that has full method set of reader interface
func retrieve(r reader) error {
	data := make([]byte, 100)
	len, err := r.read(data)

	if err != nil {
		return err
	}
	fmt.Println(string(data[:len]))
	return nil
}

type user struct {
	name string
	email string
}

type printer interface {
	print()
}

func (u user) print() {
	fmt.Printf("Sending email to %s\n", u.email)
}

func main() {

	u := user{"Adam", "adam@abc.com"}
	data := []printer{
		u,
		&u,
	}

	u.email = "John@abc.com"

	// Value semantic
	for _, e := range data {
		e.print()
	}

	f := file{"data.json"}
	p := pipe{"data.pipe"}

	retrieve(f)
	retrieve(p)
}
