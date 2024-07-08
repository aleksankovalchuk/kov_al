package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число")
	fmt.Scanln(&a)
	if a%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if a%3 == 0 {
		fmt.Println("Fizz")
	} else if a%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(a)
	}
}
