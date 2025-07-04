package main;

import("fmt")

func main() {
	fmt.Println("program to compare if two strings are anagram of each other, I.e to check if they contain same character")
	
	firstString := "ctake"
	secondString := "kacet"
	
	if len(firstString) == len(secondString) {
		controlCount := 0
		//continue your program from here
		for index := 0; index < len(firstString); index++ {
			currentRune := firstString[index]
			for indexTwo := 0; indexTwo < len(firstString); indexTwo++ {
				if currentRune == secondString[indexTwo] {
					controlCount += 1
				}
			
			}	
		}

		compareStringsLength(controlCount, firstString)

				
	}else {
		print("they cannot be compared because they are not of the same length")
	}
	//fmt.Println(strings.Contains(firstString, secondString))
}


func compareStringsLength(controlCount int, theString string) {
	if controlCount == len(theString) {	
			print("they contain same letters")
			fmt.Printf("value of control count is %d\n", controlCount)
		}else {
			print("they do not contain same letter")
			fmt.Printf("value of control count is %d\n", controlCount)
		}

}

func print(text string) {
	fmt.Println(text)
}