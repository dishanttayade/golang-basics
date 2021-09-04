package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"suitcase", "clothes", "watch"}
	fmt.Println(things)

	things = append(things, "purse")
	fmt.Println(things)

	things = append(things[1:]) //append(things[1:len(things)-1])
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "ironman"
	heros[2] = "spiderman"
	fmt.Println(heros)

	heros = append(heros, "deadpool")
	fmt.Println(heros)
	fmt.Println(cap(heros))

	myintegers := []int{4, 7, 3, 2, 8}
	isSorted := sort.IntsAreSorted(myintegers)
	fmt.Println("Are integers sorted? : ", isSorted)

	sort.Ints(myintegers)
	fmt.Println(myintegers)
}
