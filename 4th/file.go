package main

import (
	"fmt"
)

func main() {
	var isLogged bool = true
	var bal int = 10

	if isLogged && bal > 5 {
		fmt.Println("Show account info")
		var len, err = fmt.Println("Show account info")
		if err == nil {
			fmt.Printf("length is %v , %T", len, len)
			fmt.Println(`, here is the backtick`)
		}
	} else {
		fmt.Println("Show user login page")
	}
}
