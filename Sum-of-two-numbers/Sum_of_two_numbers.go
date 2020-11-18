package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Enter 1st number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter 2nd number: ")
	fmt.Scanln(&b)
	var c int = a + b
	fmt.Println(c)
}
