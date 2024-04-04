package piscine

func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	MyCount := 5
	temp := 0
	for u := 0; u < MyCount-1; u++ {
		temp = u
		for j := u + 1; j < MyCount; j++ {
			if array[j] < array[temp] {
				temp = j
			}
		}
		array[u], array[temp] = array[temp], array[u]
	}
	return array[2]
}
