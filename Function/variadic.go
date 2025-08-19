package main

import (
	"fmt"
)

func printSuppliers(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product: ", product, "Supplier: ", supplier)
	}
}
