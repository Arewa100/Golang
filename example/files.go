package main;

import("os"
	"errors"
	"log"
	)

func main() {
	file, err := os.Open("C:/Users/DELL/Desktop/Files/Design Thinking/Olasoyin Miracle.pdf")

	if err != nil {
		err = errors.New("error opening file");
		log.Fatal(err)
		
	}

	file.Close()
}	