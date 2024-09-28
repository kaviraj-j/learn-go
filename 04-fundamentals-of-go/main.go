package main

import "fmt"

func main() {

	// Variable Declaration
	// var favouriteNumber = 2
	// var myName = "Kaviraj"
	// var universityCgpa = 8.8
	// fmt.Println(favouriteNumber, myName, universityCgpa)

	// Declaring variables using short declaration operator
	favouriteNumber := 2
	myName, _, universityCgpa := "Kaviraj", 0, 8.8
	fmt.Println(favouriteNumber, myName, universityCgpa)

}
