package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	MyEvnMsg := "I have an even number of arguments"
	MyOddMsg := "I have an odd number of arguments"

	Myargs := os.Args[1:]
	LOfArgs := len(Myargs)

	if isEven(LOfArgs) {
		printStr(MyEvnMsg)
	} else {
		printStr(MyOddMsg)
	}
}