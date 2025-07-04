package main;

import("fmt")

func main() {
	fmt.Println("welcome to golang today")
	fmt.Println("we are texting booleans")
	lie := false
	truth := true
	fmt.Println(lie == lie)
	fmt.Println(!truth == lie)

	var message string = "working out"
	var messageTwo string = "workingout"
	if message == messageTwo {
		fmt.Println("this is able to compare the strings byte by byte")
	}else {
		fmt.Println("this is not able to compare the string byte by byte")
	}


	location := "77\fsopeyemi"
	fmt.Println(location)
}