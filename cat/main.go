package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func displayContent(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		printError(fileName)
		return
	}

	for _, r := range string(file) {
		z01.PrintRune(r)
	}
}

func printError(fileName string) {
	start := "ERROR: open "
	end := ": no such file or directory\n"
	ErrMessage := start + fileName + end

	for _, r := range ErrMessage {
		z01.PrintRune(r)
	}
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				break
			}
			if n == 0 {
				break
			}
			os.Stdout.Write(buf[:n])
		}
	} else if len(os.Args) == 2 {
		fileName := os.Args[1]
		displayContent(fileName)
	} else if len(os.Args) == 3 {
		fileName1 := os.Args[1]
		fileName2 := os.Args[2]
		displayContent(fileName1)
		displayContent(fileName2)
	}
}
