package piscine

func AppendRange(min, max int) []int {
	var Myappend []int
	for q := min; q < max; q++ {
		Myappend = append(Myappend, q)
	}
	return Myappend
}
