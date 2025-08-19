package main

import "fmt"

func printPrice_1(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func printPrice_2(string, float64, float64) string {
	return "Taking no parameters"
}




