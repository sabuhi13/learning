package main

import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Print(10 === 10) error
	fmt.Print(10 == 10, "\n")

	var key int = 13
	fmt.Println("key is:", reflect.TypeOf(key))

	var castedKey string = string(key)

	fmt.Println("key is:", reflect.TypeOf(castedKey))

	if 10 > 9 {
		fmt.Println("true blet!")
	}

	var getName string

	fmt.Print("enter your name: ")
	fmt.Scan(&getName)

	if getName == "sabuhi" {
		fmt.Println("yeah, baby!")
	} else if getName == "saboohy" {
		fmt.Println("qaqam coooool!")
	} else {
		fmt.Println("OMG blet!")
	}

	var myAge int = 29

	switch myAge {
		case 25:
			fmt.Println("hele cavansan")
		case 26:
			fmt.Println("evlen blet")
		case 27:
			fmt.Println("qojalersan")
		case 28:
			fmt.Println("uy blyu! qoja!")
		default:
			fmt.Println("day hec ne...")
	}
}