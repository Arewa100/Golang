package main

import ("fmt") 

const boilingPoint = 212.0

func main() {
	var f = boilingPoint
	fmt.Printf("boiling point is either %g degree F or %g degree celcius\n", f, fToC(boilingPoint))
		
	var firstNumber, secondNumber = 20, 30
	fmt.Printf("the sum of the values is %d\n", add(firstNumber, secondNumber))

	var name, age, height = "miracle", 23, 2.4 

	fmt.Printf("name: %s\nage: %d\nheight: %g\n", name, age, height)
	 
}



func fToC(givenTemp float64) float64 {
	return (givenTemp- 32) * 5 / 9    // all this are package level declaration 
}

func add(firstNumber int, secondNumber int) int {
	var sum = firstNumber + secondNumber
	return sum
}






















