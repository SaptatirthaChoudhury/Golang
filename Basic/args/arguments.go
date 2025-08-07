package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	fmt.Println("Total arguments:", len(arguments))

	for i, arg := range arguments {
		fmt.Printf("arguments[%d] = %s\n", i, arg)
	}
}

