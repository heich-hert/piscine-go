package piscine

func IsAlpha(s string) bool {
	for _, myLtr := range s {
		if !((myLtr <= 'Z' && myLtr >= 'A') || (myLtr <= 'z' && myLtr >= 'a') || (myLtr <= '9' && myLtr >= '0')) {
			return false
		}
	}
	return true
}
