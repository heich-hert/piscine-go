package piscine

func ActiveBits(n int) int {
	MyStr := ""
	MyIncreament := 0

	for n > 0 || n < 0 {
		MyStr += string(rune(n&1 + '0'))
		n = n >> 1
	}
	for _, z := range MyStr {
		if z == '1' {
			MyIncreament++
		}
	}
	return MyIncreament
}
