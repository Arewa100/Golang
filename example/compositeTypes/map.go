package main;

import("fmt")

func main() {
	//CREATING A MAP
	people := map[string]int{
			"Ayo": 20,
			"Miracle": 30,
			"Jackson": 21,
	}
	delete(people, "Jackson")
	fmt.Println(people)

	listOfNames := make(map[string]int)
	listOfNames["Ay"] = 20
	listOfNames["Shola"] = 21

	listOfNames["Shola"] += 1

	for key, value := range listOfNames {
		fmt.Printf("%s: %d\n", key, value)
	}

	//fmt.Println(listOfNames)

	//fmt.Println(len(listOfNames))


	var ages map[string]int
	
	ages["ope"] = 40

	fmt.Println(ages)
}