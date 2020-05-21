package main

/*
Pointers: Everything in Go is passed by value.

When a go program starts up, its given a logical processor
for every core available (P). And that P gets a OS thread m (memory)
scheduled by the operating system. it is also given a Go Routine, say G

There are three areas of memory for each thread: Data, stacks(Generally 1M), heap.
Similary G also gets a stack of 2k. Job of G is to execute instruction one after
another. Now G starts execution from the main function. Its gonna take a frame of
memory from stack. G has only access to memory for the frame which it is executing
from - That is the Active Frame.
The program grows down the stack. Any memory below the active frame is invalid.

Language Mechanics talk about how something works, whether the Language semantics
talk about how things behave.

Escape Analysis - this is done by the compiler to detect if an object will be pointed
down the active frame, if thats the case, it is created in the heap instead of stack.
Checks if the value escapes from stack to heap (allocation). Now GC has to get involved.

Go Stack is 2K.
What happens when it runs out of that?
it will allocate a new stack (25% more size) and copy over the stack to the new stack.

This means the stack can move.

*/

type user struct {
	name  string
	email string
}

func main() {

	// Declare a variable, count to hold an int value of 10
	count := 10
	// This happens in the main active frame

	// Lets display the value of count and its address
	println("count has a Value: ", count, " located At Addr ", &count)

	increment(count)
	// For every new function that is called, A new frame is
	// allocated and the routine starts operating in that frame
	// These frames provide a level of isolation, a kind of sandbox
	// environment to restrict its scope. This is also the reason of
	// having parameters for the function. (Mechanics)
	// This creates "program boundaries"
	/*
		The value of count of copied over to this frame as a parameter
		"Pass by Value" - Value semantics, which has benefit of
		isolation and mutability
	*/

	// Lets display the value of count and its address again after calling increment
	// It should not update since main is operating in its own isolated frame
	println("count has a Value: ", count, " located At Addr ", &count)

	/* We need Pointer semantics to share data between the boundaries
	Addresses are data! We need to store the address of the variable
	- Pointers are literal types

	Cost: Side effects(mutate) outside your isolated scope
	*/
	// Lets now pass pointer
	incrementPnt(&count)
	println("count has a Value: ", count, " located At Addr ", &count)

	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
}

// We dont have constructors in Go, but instead factories
// Create the object, initializes it

// In this version, user object is created and a copy of it returned to caller.
// whereas indirection is used to print using println
// This is valid since it is refering up the stack
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@abc.com",
	}

	println("V1", &u) //& - sharing
	return u
}

// This one uses pointer semantics
// here instead of returning a copy of the object, it is returning the address
// of it from a frame below the stack. This could be wiped out by a future call
func createUserV2() *user {
	u := user{
		name:  "john",
		email: "john@abc.com",
	}

	println("V2", &u)
	return &u
}

func increment(inc int) {

	// increment the "value of " inc
	inc++
	println("inc - Value: ", inc, " located at Addr: ", &inc)
}

func incrementPnt(inc *int) {

	// increment the "value of " inc that the "pointer points to"
	// This is indirect read/write to a memory - indirection
	*inc++
	println("inc - Value: ", inc, " located at Addr: ", &inc)
}

/*
$ go build -gcflags "-m -m" pntr.go
# command-line-arguments
./pntr.go:105:6: can inline increment as: func(int) { inc++; println("inc - Value: ", inc, " located at Addr: ", &inc) }
./pntr.go:112:6: can inline incrementPnt as: func(*int) { *inc++; println("inc - Value: ", inc, " located at Addr: ", &inc) }
./pntr.go:79:6: can inline createUserV1 as: func() user { u := user literal; println("V1", &u); return u }
./pntr.go:93:6: can inline createUserV2 as: func() *user { u := user literal; println("V2", &u); return &u }
./pntr.go:31:6: cannot inline main: function too complex: cost 99 exceeds budget 80
./pntr.go:40:11: inlining call to increment func(int) { inc++; println("inc - Value: ", inc, " located at Addr: ", &inc) }
./pntr.go:64:14: inlining call to incrementPnt func(*int) { *inc++; println("inc - Value: ", inc, " located at Addr: ", &inc) }
./pntr.go:67:20: inlining call to createUserV1 func() user { u := user literal; println("V1", &u); return u }
./pntr.go:68:20: inlining call to createUserV2 func() *user { u := user literal; println("V2", &u); return &u }
./pntr.go:94:2: u escapes to heap:
./pntr.go:94:2:   flow: ~r0 = &u:
./pntr.go:94:2:     from &u (address-of) at ./pntr.go:100:9
./pntr.go:94:2:     from return &u (return) at ./pntr.go:100:2
./pntr.go:94:2: moved to heap: u
./pntr.go:112:19: inc does not escape
*/

/*
Go uses tri-color mark and sweep concurrent collector.

First, all the objects are marked white and then root objects are marked black
and if the object is child of another, that is marked grey. In subsequent rounds
if grey objects are pointing to another objects, they are marked black and
so on.

Goals are to keep heap as small as possible, clear the heap at less than 100ms
and not take more than 25% of the CPU time in GC.
*/
