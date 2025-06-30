package main;

import("fmt")

func main() {

	//pointers variables
	mainValue := 2;
	fmt.Println(mainValue);
	mainValueAddress  := &mainValue;
	fmt.Println(*mainValueAddress);

	*mainValueAddress = 24;

	fmt.Printf("this is the value assigned through the pointer %d \n", mainValue)
		
	testValue  := 4;
	
	secondPointer := &testValue;
	
	if(mainValueAddress == secondPointer) {
		fmt.Println("success")
	}else {
		fmt.Printf("error, the pointers are pointing %d, %d to different memory address", mainValueAddress, secondPointers);
	}
}
	