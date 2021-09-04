package main

import (
	"fmt"
)

func main() {
	// make new
	// new - only allocates - no init of memory
	// make - allocate and init - non zeroed storage

	// var score map[string]int
	score := make(map[string]int)
	score["first"] = 89
	fmt.Println(score)

	score["second"] = 54
	score["third"] = 78
	// score["second"] = 58
	fmt.Println(score)

	delete(score, "third")
	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}
