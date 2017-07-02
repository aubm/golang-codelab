package basics

// Pay attention to the name of this function.
// You do see that it is lying to you right?
// You will definitely not let that happen!
func AlwaysInvokeCallbackBeforeExitButOnlyWriteItOnce(someNumber int, callback func(int)) {
	if someNumber < 0 {
		someNumber = 0
		callback(someNumber)
		return
	}

	if someNumber < 10 {
		someNumber = 10
		callback(someNumber)
		return
	}

	if someNumber < 50 {
		someNumber = 50
		callback(someNumber)
		return
	}

	someNumber = 100
	callback(someNumber)
}
