package main;
import("fmt")

func main() {

	fmt.Println("welcome to immutable byte")
	name := "Miracle"
	fmt.Println(name[0:4])
	fmt.Println(name[0:4] + "bel")
	for index := 0; index < len(name); index++ {
		fmt.Println(name[index])
	}
}