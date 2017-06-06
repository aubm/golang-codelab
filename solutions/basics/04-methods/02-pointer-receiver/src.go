package basics

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetOlder() {
	p.Age++
}
