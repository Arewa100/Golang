package main;

import("fmt") 


func main() {
	fmt.Println("welcome to the pointer program");
	
	//fmt.Println(store());
	//fmt.Println(store());

	theValue := 1;

	result := firstIncrementor(theValue)
	fmt.Printf("the result of the function is %d and the original value is now %d without using the pointer method \n\n ", result, theValue)

	theSecondValue := 1;
	secondResult := secondIncrementor(&theSecondValue)
	fmt.Printf("the current value of thesecondvalue variable is %d, and increment is %d", theSecondValue, secondResult);
}

func store() *int {
	number := 2;
	return &number;
}


//i want to modify a variable without referencing it address

func firstIncrementor(value int) int {
	value = value + 1;
	return value;
}

func secondIncrementor(value *int) int {
	*value = *value + 1;
	return *value;
}