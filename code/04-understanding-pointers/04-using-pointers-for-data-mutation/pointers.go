package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int = &age

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Age after editAgeToAdultYears:", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
