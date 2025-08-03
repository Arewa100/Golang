package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Harte Airlines")
	var numberOfSeats [10]bool = [10]bool{}
	var count string = "yes"
	for count != "no" && count == "yes" {
		fmt.Println("select 1 for first class and 2 for economy class")
		var userSelection int
		selection, err := fmt.Scan(&userSelection)
		if err != nil {
			log.Fatal(errors.New("error reading user selection"), selection)

		}

		if userSelection == 1 {
			var sitNumber int
			display := true
			for index := 0; index < 5; index++ {
				if numberOfSeats[index] != true {
					numberOfSeats[index] = true
					sitNumber = index + 1
					break
				} else if index == 4 && numberOfSeats[index] == true {
					fmt.Println("first class is occupied, do you want to be transferred to economy class? press yes to continue and no to end application")
					scan, err := fmt.Scan(&count)
					if err != nil {
						log.Fatal(errors.New("error reading user selection"), scan)
					}
					display = false
				}
			}

			if display != false {
				response := "your sit number is" + " " + strconv.Itoa(sitNumber) + " " + "and" + " " + "you are in first class"
				fmt.Println(response)
			}
		}

		if userSelection == 2 {
			var sitNumber int
			display := true
			for index := 5; index < 10; index++ {
				if numberOfSeats[index] != true {
					numberOfSeats[index] = true
					sitNumber = index + 1
					break
				} else if index == 9 && numberOfSeats[index] == true {
					fmt.Println("economy class is occupied, do you want to be transferred to first class? press yes to continue and no to end application")
					scan, err := fmt.Scan(&count)
					if err != nil {
						log.Fatal(errors.New("error reading user selection"), scan)
					}
					display = false
				}
			}
			if display != false {
				response := "your sit number is" + " " + strconv.Itoa(sitNumber) + " " + "and" + " " + "you are in economy class"
				fmt.Println(response)
			}
		}

	}
}
