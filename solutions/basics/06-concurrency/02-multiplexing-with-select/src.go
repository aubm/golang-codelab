package basics

var (
	Margarita               = make(chan int)
	Manhattan               = make(chan int)
	Wine                    = make(chan int)
	Biers                   = make(chan int)
	CanWeHaveTheCheckPlease = make(chan bool)
)

type Slate struct {
	Biers, Margarita, Manhattan, Wine int
}

// Because we don't want to use our credit card each time we order a drink,
// we will start a slate for the night, and will only pay the note once
// at the end of the night.
// Using the global channels defined above, we can order drinks, and ask for
// the check at the end of the night.
// When calling StartSlate, the function must block until a value is published
// in CanWeHaveTheCheckPlease.
// In the meantime, each time a certain quantity of a drink is ordered, using
// one of the Margarita, Manhattan, Wine or Biers channels, that quantity must
// be added to the current slate.
// A value will be published in CanWeHaveTheCheckPlease only once. When that happens,
// StartSlate must return the slate.
func StartSlate() (slate Slate) {
	for {
		select {
		case n := <-Biers:
			slate.Biers += n
		case n := <-Margarita:
			slate.Margarita += n
		case n := <-Manhattan:
			slate.Manhattan += n
		case n := <-Wine:
			slate.Wine += n
		case <-CanWeHaveTheCheckPlease:
			return
		}
	}
}
