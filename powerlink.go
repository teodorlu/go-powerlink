package powerlink

type Person struct {
	name   string
	father *Person
	mother *Person
}

// do I need generics in order to get typed access?
type Powerlink interface {
	// ....
	PowerlinkRefer()
}

// what are the types?
//
// I don't want to use structs to navigate.
// I want to use rich relations.
// Typed relations.
//
// Where one can /statically link/ by referring to the same link, yet /not/ use struct access.
//
// Which can enable us to navigate to links based on relation semantics --

type FatherRel struct {
	from Person
	to   Person
}

// Now I'd like to traverse a link with /link semantics/ and control the /return type/.

func PowerlinkNav() {

}

// How do I set the return type right here?
// Return type depends on the powerlink.
func Get(x any, l Powerlink) {}

// This one is even harder to type.
// But I guess I can skip it.
func GetIn(x any, chain []Powerlink) {}
