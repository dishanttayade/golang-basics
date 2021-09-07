package main

import (
	"fmt"
)

func main() {
	var temp int = 5

	switch temp {
	case 1:
		fmt.Println("One")
	case 2, 3:
		fmt.Println("Two or three")
	case 4:
		fmt.Println("Four")
		// fallthrough
		// to continue to next case
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("I don't know")
	}

	// For comparison
	tem := -5
	switch {
	case tem < 0:
		fmt.Println("It's cold")
	case tem == 0:
		fmt.Print("Cool, huh?")
	case tem > 0:
		fmt.Println("Yep, Its hot")
	}
}
