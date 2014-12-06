// leap contains functions to compute leap years
package leap

// Isleap returns true if the argument is a leap year.
func Isleap(y int) bool {
	if y%400 == 0 {
		return true
	}
	if y%100 == 0 {
		return false
	}
	if y%4 == 0 {
		return true
	}
	return false
}


