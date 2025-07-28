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
	fmt.Println(listOfNames)
}