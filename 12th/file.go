package main

import (
	"fmt"
)

func main() {
	afunction()
	re := multiply(3, 6)
	fmt.Println(re)

	res, len, name := add(-20, -6, 5, 8, 10, 11, 13, 19, 22, 25, 27, 30)
	fmt.Println(res, len, name)

}

func afunction() {
	fmt.Println("Hey I am a function")
}

func multiply(a int, b int) int {
	return a * b

func add(vals ...int) (int, int, string) {
	sum := 0
	for i := range vals {
		sum = sum + vals[i]
	}
	len := len(vals)
	name := "fun"
	return sum, len, name
}
