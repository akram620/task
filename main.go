package main

import (
	"fmt"
	"task/solve"
)

func main() {
	// TotalAmount with parameters: productName, amount, phoneNumber, duration
	fmt.Println("Total amount:", solve.TotalAmount("phone", 1000, "+992 93 988 88 88", 18))
}
