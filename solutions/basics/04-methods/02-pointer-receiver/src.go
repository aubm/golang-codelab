package basics

// When calling GetOlder on a given Person, it does not actually add a year to that person.
// Any idea about what should be done to fix it?

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetOlder() {
	p.Age++
}
