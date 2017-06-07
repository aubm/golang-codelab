package basics

func GiveMeAPersonAndIWillRenameItToJohnDoe(person *Person) {
	person.FirstName = "John"
	person.LastName = "Doe"
}

type Person struct {
	FirstName string
	LastName  string
}
