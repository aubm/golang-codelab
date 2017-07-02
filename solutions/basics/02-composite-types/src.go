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

// When given ["My", "favorite", "language", "is", "Golang"], it should return [2, 8, 8, 2, 6]
func MustCountTheNumberOfLettersInEachWord(words []string) []int {
	count := []int{}
	for _, word := range words {
		count = append(count, len(word))
	}
	return count
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

// The function must panic if the json is invalid
// Learn more about the panic built-in here: https://gobyexample.com/panic
// You might want to read about Go and JSON: https://gobyexample.com/json
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
