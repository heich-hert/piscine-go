package piscine

func AlphaCount(s string) int {
	WordNumber := 0

	for _, lt := range s {
		if (lt <= 'z' && lt >= 'a') || (lt <= 'Z' && lt >= 'A') {
			WordNumber++
		}
	}

	return WordNumber
}
