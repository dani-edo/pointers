package main

import (
	"fmt"
)

func main() {
	age := 32 // regular variable
	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println("Adult years:", adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}