package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------------------")

	number := 1

	for number <= 5 {
		fmt.Println(number)

		number++
	}
}