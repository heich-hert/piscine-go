package piscine

func IsUpper(s string) bool {
	for _, myLt := range s {
		if myLt > 'Z' || myLt < 'A' {
			return false
		}
	}
	return true
}
