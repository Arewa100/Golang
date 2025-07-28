package main;

import("fmt")

func main() {
	fmt.Println("welcome to slices")

	mySlice := make([]byte, 4, 7)
	runes := append(mySlice, 77)
	rune := append(mySlice, 55)
	fmt.Println(runes)
	fmt.Println(rune)
	fmt.Println(mySlice)
}