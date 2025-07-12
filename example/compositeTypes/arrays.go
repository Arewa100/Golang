package main;

import("fmt")

func main() {
	//composite types in go
	fmt.Println("composite is working")

	var myByteArray [3]byte = [3]byte{77, 106, 100}
	fmt.Println(myByteArray)

	//for index, element := range myByteArray {
		//fmt.Printf("%d %d\n", index, element)
	//}
	for index := len(myByteArray)-1; index >= 0; index-- {
		element := myByteArray[index]
		//valueInBinary := strconv.FormatInt(int64(element), 2)
		fmt.Printf("%d \n", element)
	}


	list := [...]string{ "ope", "ayo", "gbemileke"}
	fmt.Println(list)
	
}