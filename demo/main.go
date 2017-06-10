package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	someRandomName := randomdata.FullName(0)
	fmt.Println("Hello " + someRandomName)
}
