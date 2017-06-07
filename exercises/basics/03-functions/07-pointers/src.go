package basics

// This function does not do what it say it does.
// Any idea about what should be done in order to fix it?
func GiveMeAPersonAndIWillRenameItToJohnDoe(person Person) {
	person.FirstName = "John"
	person.LastName = "Doe"
}

type Person struct {
	FirstName string
	LastName  string
}
