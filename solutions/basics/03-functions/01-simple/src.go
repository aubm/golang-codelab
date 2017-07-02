package basics

import "fmt"

func rgbToHex(red, green, blue int) string {
	return fmt.Sprintf("%02X%02X%02X", red, green, blue)
}

// In order for the tests to pass in this package, you must implement the three following functions:
// Red, Green and Blue.
// Red, Green and Blue are functions that return a string, which are the color hex for the colors red, green and blue.
// Get here if you need more information about colors hex: https://www.w3schools.com/colors/colors_hexadecimal.asp
// You should use the provided rgxToHex helper function.

func Red() string {
	return rgbToHex(255, 0, 0)
}

func Green() string {
	return rgbToHex(0, 255, 0)
}

func Blue() string {
	return rgbToHex(0, 0, 255)
}
