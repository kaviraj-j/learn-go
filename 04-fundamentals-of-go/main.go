package main

import "fmt"

func main() {

	// Variable Declaration
	// var favouriteNumber = 2
	// var myName = "Kaviraj"
	// var universityCgpa = 8.8
	// fmt.Println(favouriteNumber, myName, universityCgpa)

	// Declaring variables using short declaration operator
	/*
	 */

	favouriteNumber := 2
	myName, _, universityCgpa := "Kaviraj", 0, 8.8
	fmt.Println(favouriteNumber, myName, universityCgpa)

	fmt.Println("************** Conversion In go **************")
	intNum := 30
	floatNum := 23.92232
	fmt.Printf("%v of type %T \n", intNum, intNum)
	fmt.Printf("%v of type %T \n", floatNum, floatNum)
	intNum = int(floatNum)
	fmt.Println(intNum)

}
