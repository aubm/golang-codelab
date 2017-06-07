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
