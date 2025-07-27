package main;

import("fmt")

type Currency int

const(
	USD Currency = iota
	NGN	
	AUD
)

func main() {
	fmt.Println("my arrays")
	listOfBooks := [...]string{}
	fmt.Printf("%T\n", listOfBooks)

	symbols := [...]string{USD: "$", NGN: "n", AUD: "au"}

	fmt.Println(symbols[2])


	testArray := [4]int{2, 3, 4, 6}
	fmt.Println(testArray)
	pointer := &testArray
	manipulateArrayData(pointer)
	fmt.Println(testArray)
	
	
}

func manipulateArrayData(array *[4]int) {
	array[2] = 10
}

