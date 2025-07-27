package main;

import("fmt")

func main() {
	fmt.Println("welcome to the square testing")

	var number int
	pointer := &number //*int
	fmt.Println("enter a number")
	fmt.Scan(pointer)
	
	square, cube := squareAndCube(pointer)
	fmt.Printf("the square and cube of the number is %d %d", square, cube)
	
	
	
}

func squareAndCube(number *int) (square int, cube int) {
	square = *number * *number
	cube = *number * *number * *number
	return square, cube
}