package basics

import "fmt"

// The Meeting.GetStarted method will not compile because the Meeting does not have
// a Introducer property yet.
// Can you fix it?
// Meeting.Introducer should accept any of types Capuchin, Baboon or Mandrill as a value.
// Also, Meeting.Introducer should not have a named type.

type Meeting struct {
}

func (m Meeting) GetStarted() {
	m.Introducer.Introduce()
}

type Capuchin struct{}

func (c Capuchin) Introduce() {
	fmt.Println("I am a capuchin!")
}

type Baboon struct{}

func (b Baboon) Introduce() {
	fmt.Println("I am a baboon!")
}

type Mandrill struct{}

func (m Mandrill) Introduce() {
	fmt.Println("I am a mandrill!")
}
