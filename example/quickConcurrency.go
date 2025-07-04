package main;
import("fmt"
	"os"
	"io")

func main() {

	//go printName();
	//fmt.Println("this is the second statement")

	//declaring a variable using keyword new

	//first := new(int)
	//*first = 4
	//fmt.Println(*first)

	file, err := os.Open("C:/Uers/DELL/Desktop/Files/helloGo/example/pointers.go")

	if err != nil {
		fmt.Errorf("an error occured %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading the file %", err)
	}
	fmt.Printf("content is readable \n%s\n", string(content))
	
}


func printName() {
	fmt.Println("My name is Miracle");
}
