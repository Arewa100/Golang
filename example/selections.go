package main;
import("fmt"
	"errors")

func main() {
	result, err := divide(3, 3)
	fmt.Println(result)
	fmt.Println(err)
}


func divide(first float64, second float64) (result float64, err error) {
	if(second == 0) {
		err = errors.New("cannot divide by zero")
		return 
	}
	result = first/second;
	return result, err
	  
}