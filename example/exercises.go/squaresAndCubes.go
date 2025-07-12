package main;

import("fmt")

func main() {
	fmt.Println("welcome to the square printer")
	fmt.Println("")
	var number int
	fmt.Println("enter an integer: ")
	fmt.Scanf("%d\n", &number)

	//using a while loop
	fmt.Printf("%-10s %-10s %s\n", "numbers", "square", "cube")
	count := 0
	for count <= number {
		square := count * count
		cube := count * square
		fmt.Printf("%-10d %-10d %d\n", count, square, cube) 
		count++
	}
	 
}