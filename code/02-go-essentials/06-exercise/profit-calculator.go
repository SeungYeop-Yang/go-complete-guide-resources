package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue Amount: ")
	fmt.Scan(&revenue)

	fmt.Print("expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("tax rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / profit

	fmt.Println(earningBeforeTax, profit, ratio)
}
