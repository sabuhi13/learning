package main

import (
	"fmt"
)

// comment 1

/*
	comment 2
*/

func main() {

	/**
	 * Varianbles
	 */
	var a int = 1
	var b = 2
	c := 3

	var d int
	var e int
	e = 4
	e = 5

	fmt.Println("dot dot doooot!", a, b, c, d, e)

	var name, surname, age = "Sabuhi", "Alizada", 28

	fmt.Println(name, surname, age)

	w1, w2 := "Word 1", "Word 2"

	fmt.Println(w1, w2)

	const c1 int = 2023
	const c2 = 2050
	fmt.Println("Constants:", c1, c2)

	var id, count, year int
	id = 1
	count = 13
	year = 2023
	
	fmt.Println(id, count, year)

	/**
	 * Data types
	 */
	var integerData int = 1 // 32/64
	var unsignetIntegerData uint = 2 // 32/64

	fmt.Println(integerData, unsignetIntegerData)

	var c3 int8 = 16
	var c4 uint8 = 17
	fmt.Println(c3, c4)

	var c5 int16 = 128
	var c6 uint16 = 128
	fmt.Println(c5, c6)

	var c7 int32 = 256
	var c8 uint32 = 256
	fmt.Println(c7, c8)

	var c9 int64 = 1024
	var c10 uint64 = 1024
	fmt.Println(c9, c10)

	var salary32 float32 = 10.11
	var salary64 float64 = 10.13
	salary := 10.14 // 64 bit float
	fmt.Println(salary32, salary64, salary)

	var myName string = "Sabuhi"
	var mySurname string = "Alizada"
	fmt.Println(myName, mySurname)

	var isValid bool = true
	fmt.Println(isValid)

	/**
	 * Print statement
	 */
	fmt.Print("yeni")
	fmt.Print("bir")
	fmt.Print("nese\n")

	testValue := "test value"
	fmt.Print("printed: ", testValue, "\n")

	fmt.Println("yeni bir nese")

	var myAge int8 = 28
	fmt.Printf("My Age: %d\n", myAge)

	var v1, v2, v3, v4 = 10, 10.13, "value", true
	fmt.Printf("integer: %d, float: %g, string: %s, boolean: %t\n", v1, v2, v3, v4)

	print("hello\n")
	println("hello")
	// testName := "sabuhi"
	// printf("hello, %s", testName) error

	/**
	 * Input
	 */
	var getName string

	fmt.Print("Enter your name: ")
	fmt.Scan(&getName) // Reference is required

	fmt.Println("Your name is", getName)

	var inputA, inputB string
	fmt.Println("Enter anyting: ")
	fmt.Scan(&inputA, &inputB)

	fmt.Printf("values: %s %s\n", inputA, inputB)

	var inputC, inputD string
	fmt.Println("Enter anything again:")
	fmt.Scanln(&inputC, &inputD) // write two values

	fmt.Printf("values: %s %s\n", inputC, inputD)

	var inputE, inputF string
	fmt.Println("Enter anything again:")
	fmt.Scanf("%s %s", &inputE, &inputF) // write two values

	fmt.Printf("values: %s %s\n", inputE, inputF)
}