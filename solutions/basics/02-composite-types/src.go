package basics

import (
	"encoding/json"
	"sort"
	"strings"
)

func MustReturnAnArrayOfThreeIntegers() [3]int {
	return [3]int{}
}

func MustReturnTheFirstElementOfSlice(elements []int) int {
	return elements[0]
}

func MustReturnTheLastElementOfSlice(elements []int) int {
	return elements[len(elements)-1]
}

func MustReturnTheGreatestElementOfSlice(elements []int) int {
	var greatestElement int
	for _, element := range elements {
		if element > greatestElement {
			greatestElement = element
		}
	}
	return greatestElement
}

func MustReturnTheNumberOfWords(text string) int {
	if text == "" {
		return 0
	}
	words := strings.Split(text, " ")
	return len(words)
}

func MustAppendTheWordGolang(words []string) []string {
	return append(words, "Golang")
}

func MustReturnTheLevelKeyOrOneByDefault(stats map[string]int) int {
	level, ok := stats["level"]
	if !ok {
		level = 1
	}
	return level
}

func MustSortAndReturn(listOfIntegers []int) []int {
	sort.Ints(listOfIntegers)
	return listOfIntegers
}

func MustReturn21YearsOldJohnDoe() Person {
	return Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       21,
	}
}

// The function must panic is the json is invalid
// Learn more about the panic built-in here: https://gobyexample.com/panic
func MustReturnAListOfPerson(jsonListOfPerson []byte) []Person {
	listOfPerson := make([]Person, 0)
	if err := json.Unmarshal(jsonListOfPerson, &listOfPerson); err != nil {
		panic(err)
	}
	return listOfPerson
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int
}
