package main;

import("fmt")


func main() {
	fmt.Println("welcome to the gas milage calculator")
	//using sentinel controlled iteration 
	sentinel := 0
	var combinedMilesPerGallon float64 = 0

	for sentinel != -1 {
		var milesDriven int
		var gallonsUsed int
		fmt.Println("enter the miles driven for this trip")
		fmt.Scanf("%d\n", &milesDriven)
		fmt.Println("enter the gallons used")
		fmt.Scanf("%d\n", &gallonsUsed)
		fmt.Printf("the miles per gallon is %f\n", float64(milesDriven/gallonsUsed))
		combinedMilesPerGallon = combinedMilesPerGallon + float64(milesDriven/gallonsUsed)
		fmt.Println("click -1 to end or enter to coninue")
		fmt.Scanf("%d\n", &sentinel)
	}
}