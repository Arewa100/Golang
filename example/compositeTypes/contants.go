package main;

import("fmt")

type currency int

const (
	EUR currency = iota
	DOL
	AUD
	NG
)

func main() {
	fmt.Println("arrays and constants")
	symbols := [...]string{EUR: "e", DOL:"$", AUD:"au", NG:"#"}
	fmt.Println(symbols[0])

//this above is like three things in one, key value pair, also indexing 

}