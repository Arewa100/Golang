package main;

import("fmt"; "time")

func main() {
	
	fmt.Println(time.Now()) 
	
	const NUMBER = 24;
	fmt.Println(NUMBER);

	//declaring a set of variables
	var firstName, secondName, middleName string= "Miracle", "Olasoyin", "Talomola";
	fmt.Printf("this are ur names %s, %s, %s", firstName, secondName, middleName);
}

func Add(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber;
}

func Substract(firstNumber int, secondNumber int) int {
	return firstNumber - secondNumber;
}