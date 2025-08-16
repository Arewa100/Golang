package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Data struct {
	Name string
	Id   int
}

func main() {
	schoolData := []Data{
		{"ope", 1},
		{"jack", 2},
		{"girl", 3},
	}

	newJsonData, err := json.Marshal(&schoolData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newJsonData)

	var theConvertedData []struct {
		Name string
		Id   int
	}

	err = json.Unmarshal(newJsonData, &theConvertedData)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(theConvertedData)
	//fmt.Println(theConvertedData[0].Id)
}
