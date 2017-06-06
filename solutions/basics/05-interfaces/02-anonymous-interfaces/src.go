package basics

import "fmt"

type Meeting struct {
	Introducer interface {
		Introduce()
	}
}

func (m Meeting) GetStared() {
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
