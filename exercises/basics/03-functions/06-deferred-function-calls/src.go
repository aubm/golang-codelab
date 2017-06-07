package basics

// Pay attention to the name of this function.
// You do see that it is lying to you right?
// You will definitely not let that happen!
func AlwaysInvokeCallbackBeforeExitButOnlyWriteItOnce(someNumber int, callback func()) {
	if someNumber < 0 {
		callback()
		return
	}

	if someNumber < 10 {
		callback()
		return
	}

	if someNumber < 50 {
		callback()
		return
	}

	callback()
}
