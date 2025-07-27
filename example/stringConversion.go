package main;

import("fmt"
	"strconv"
	"errors")


func main() {
	print("welcome to the integer to string converter")
	age := "77"
	//y := fmt.Sprintf("%d", age)
	//fmt.Println(strconv.FormatInt(int64(age), 2))

	x, err := strconv.Atoi(age)
	if err != nil {
		err = errors.New("error o")
		fmt.Errorf("error o %v", err)
	}

	fmt.Println(x + 4)
}

func print(text string) {
	fmt.Println(text)
}to