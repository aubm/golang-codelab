package basics

// Pay attention to the name of this function.
// You do see that it is lying to you right?
// You will definitely not let that happen!
func AlwaysInvokeCallbackBeforeExitButOnlyWriteItOnce(someNumber int, callback func(int)) {
	defer func() {
		callback(someNumber)
	}()

	if someNumber < 0 {
		someNumber = 0
		return
	}

	if someNumber < 10 {
		someNumber = 10
		return
	}

	if someNumber < 50 {
		someNumber = 50
		return
	}

	someNumber = 100
}
