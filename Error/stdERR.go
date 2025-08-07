package main

import (
	"io"
	"os"
)

func Error() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
		io.WriteString(os.Stdout, "This is the Standard output\n")
		io.WriteString(os.Stdout, myString)
	} else {
		myString = arguments[1]
		io.WriteString(os.Stderr, myString)
		io.WriteString(os.Stderr, "\n")
	}

}
