package piscine

func BasicJoin(elems []string) string {
	FullLength := 0
	for _, elem := range elems {
		FullLength += len(elem)
	}
	MyOutput := make([]byte, 0, FullLength)
	for _, Mylmnt := range elems {
		MyOutput = append(MyOutput, []byte(Mylmnt)...)
	}
	return string(MyOutput)
}
