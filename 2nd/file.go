package main

import "fmt"

func main() {
	var str string = "Hey there!"
	fmt.Println(str)

	var another string
	another = "Hey again"
	fmt.Println(another)

	onemore := "Hey, you again"
	fmt.Println(onemore)

	num := 3
	// num := 3.
	// num := 3.0
	fmt.Printf("%v, %T", num, num)

	var One, two string = "one", "two"
	fmt.Println(One, two)

	// var default int
	var defaultVal string
	fmt.Println(defaultVal)
}
