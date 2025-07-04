package main;
import("fmt")

func main() {
	name := "opeyemi"
	fmt.Println(&name)
	bytename := []byte(name)
	fmt.Printf("this is the address of the byte %p\n", &bytename) 

	bytename[0] = 77
	fmt.Println(bytename)

	name = string(bytename)
	fmt.Println(name)
	fmt.Print(&name)
}