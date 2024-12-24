package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int

	agePointer = &age

	fmt.Println("AgePointer:", agePointer)
	fmt.Println("Age:", *agePointer)

	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
