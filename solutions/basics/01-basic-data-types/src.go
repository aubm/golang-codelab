package basics

func MustReturnAnInteger() int {
	return 0
}

func MustReturnTheSum(a, b int) int {
	return a + b
}

func MustReturnAFloat() float64 {
	return 1.0
}

func MustBeTrue() bool {
	return true
}

func MustBeFalse() bool {
	return false
}

func MustBeTrueIfThereAreTenEggsOrMore(numberOfEggs int) bool {
	return numberOfEggs >= 10
}

func MustReturnHello() string {
	return "Hello"
}

// Here you should only add new lines without changing the existing ones
func MustReturnWorld() string {
	who := "world"
	return who
}

// Here you should only add new lines without changing the existing ones
func MustSetToTrueAndReturn() bool {
	var changeMe bool
	changeMe = true
	return changeMe
}
