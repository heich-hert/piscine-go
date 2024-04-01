package main

import "github.com/01-edu/z01"

func main() {
	for al := 'z'; al >= 'a'; al-- {
		z01.PrintRune(al)
	}

	z01.PrintRune('\n')
}
