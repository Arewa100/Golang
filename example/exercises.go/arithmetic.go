package main;

import("fmt")

func main() {
	fmt.Println("welcome to arithmetic\n")
	var first int
	var second int
	fmt.Println("enter first numbers")
	fmt.Scanf("%d\n", &first)
	fmt.Println("enter second numbers")
	fmt.Scanf("%d\n", &second)
	
	firstNumberSquare := first * first
	secondNumberSquare := second * second

	sumOfTheirSquare := firstNumberSquare + secondNumberSquare
	diffOfTheirSquare := firstNumberSquare - secondNumberSquare
	fmt.Printf("the sum of their square is %d\n", sumOfTheirSquare)
	fmt.Printf("the difference of their square is %d\n", diffOfTheirSquare)

	
	
}