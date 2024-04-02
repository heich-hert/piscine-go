package piscine

func IsLower(s string) bool {
	for _, Lowerlt := range s {
		if Lowerlt > 'z' || Lowerlt < 'a' {
			return false
		}
	}
	return true
}
