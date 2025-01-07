package main

import (
	"fmt"
)

func main() {
	age := 32 // regular variable
	fmt.Println("Age:", age)

	editAgeToAdultYears(&age)
	fmt.Println("Adult years:", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18 // mutate age variable value in the same pointer
}
