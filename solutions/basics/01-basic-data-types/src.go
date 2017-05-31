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
	return "hello"
}

func MustReturnWorld() string {
	who := "world"
	return who // do not edit this line
}

func MustSetToTrueAndReturn() bool {
	var changeMe bool // do not edit the line
	changeMe = true
	return changeMe // do not edit the line
}
