package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type User struct {
		Name    string
		Age     int
		Hobbies []string
	}

	//creating an array of users, wow this is an array of reference types. i am so amazed
	userArray := []User{
		{Name: "ola", Age: 20, Hobbies: []string{"eating", "singing", "swimming"}},
		{Name: "ola", Age: 20, Hobbies: []string{"eating", "singing", "swimming"}},
		{Name: "ola", Age: 20, Hobbies: []string{"eating", "singing", "swimming"}},
	}

	//now let's encode this go array into json ...that is an ordered sequence of values represented as comma seperated list
	fmt.Printf("user Array before marshalling %T\n\n", userArray)

	data, err := json.Marshal(userArray)
	if err != nil {
		panic(err)
	}

	fmt.Printf("user Array after marshalling %s\n", data)

	//unmarshalling

	var incomingData []struct {
		Name    string
		Age     int
		Hobbies []string
	}

	err = json.Unmarshal(data, &incomingData) //this is literally deserializing data
	if err != nil {
		panic(err)
	}
	fmt.Printf("user Array after unmarshalling %s\n", incomingData[0].Name)

}
