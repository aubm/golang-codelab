package basics

func MustReturnAnArrayOfThreeIntegers() [3]int {
}

func MustReturnTheFirstElementOfSlice(elements []int) int {
}

func MustReturnTheLastElementOfSlice(elements []int) int {
}

func MustReturnTheGreatestElementOfSlice(elements []int) int {
}

func MustReturnTheNumberOfWords(text string) int {
}

func MustAppendTheWordGolang(words []string) []string {
}

func MustReturnTheLevelKeyOrOneByDefault(stats map[string]int) int {
}

func MustSortAndReturn(listOfIntegers []int) []int {
}

func MustReturn21YearsOldJohnDoe() Person {
}

// The function must panic if the json is invalid
// Learn more about the panic built-in here: https://gobyexample.com/panic
func MustReturnAListOfPerson(jsonListOfPerson []byte) []Person {
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int
}
