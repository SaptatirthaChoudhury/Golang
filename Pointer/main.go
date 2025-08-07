package main

import (
	"fmt"
)

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is\n", a)

	c := 25
	var d *int
	//lint:ignore SA5011 safe to check for nil
	if d == nil {
		fmt.Println("d is", d)
		d = &c
		fmt.Println("d after initialization is", d)
	}
}
