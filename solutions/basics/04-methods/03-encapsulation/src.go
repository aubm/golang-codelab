package basics

func NewPerson(age int) *Person {
	if age < 1 || age > 120 {
		age = 1
	}
	return &Person{age: age}
}

type Person struct {
	Name string
	age  int
}

func (p *Person) GetOlder() {
	p.age++
}

func (p Person) Age() int {
	return p.age
}
