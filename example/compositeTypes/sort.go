package main;

import("fmt")

func main() {
	fmt.Println("enter five numbers")
	var arrayOfNumbers [5]int = [5]int{}
	
	for index := 0; index < 5; index++ {
		number := 0
		fmt.Println("enter a number")
		fmt.Scan(&number)
		arrayOfNumbers[index] = number
	}



	sortNumbers(arrayOfNumbers)
}


func sortNumbers(arrayOfNumbers [5]int) {
	for index := 0; index < len(arrayOfNumbers); index++ {
		currentElement := arrayOfNumbers[index]   //2
		for secondIndex := index + 1; secondIndex < len(arrayOfNumbers); secondIndex++ {
			if currentElement > arrayOfNumbers[secondIndex] {
				previousElement := currentElement
				arrayOfNumbers[index] = arrayOfNumbers[secondIndex]
				arrayOfNumbers[secondIndex] = previousElement
				currentElement = arrayOfNumbers[index]
				fmt.Println(currentElement)
			}
		}
	}

	fmt.Println(arrayOfNumbers)
}
