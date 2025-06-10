package main

import (
	"example/workspace/module1/calculator"
	"fmt"
)

func main() {
	sum := calculator.Add(5, 3)
	product := calculator.Multiply(4, 2)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}
