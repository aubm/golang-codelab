package basics

import "fmt"

func rgbToHex(red, green, blue int) string {
	return fmt.Sprintf("%02X%02X%02X", red, green, blue)
}

func Red() string {
	return rgbToHex(255, 0, 0)
}

func Green() string {
	return rgbToHex(0, 255, 0)
}

func Blue() string {
	return rgbToHex(0, 0, 255)
}
