package main

import ("fmt")

const HEIGHT = 12
var name string = "Miracle"
func main() {

	for index := 1; index <=5; index++ {
		fmt.Printf("the height is %d \n ", HEIGHT)
	}
	var result int = add(4, 5)
	fmt.Printf("the result is %d \n", result)
	
	for index := 1; index <= 3; index++ {
		print("i love Eri")
	}
	
	print(name)
	
}

func add(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func print(word string) {
	fmt.Println(word)
}