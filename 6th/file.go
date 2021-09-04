package main

import (
	"fmt"
)

func main() {
	var value int = 23

	// var p *int
	p := &value

	if p != nil {
		fmt.Println("The value of p : ", *p)
	} else {
		fmt.Println("No value")
	}

	fmt.Println(value)
	fmt.Println(p)

	// Array

	var nums [3]string
	nums[0] = "uno"
	nums[1] = "dos"
	nums[2] = "tres"
	fmt.Println(nums)

	var colors = [4]string{"rojo", "gris", "azul", "verde"}
	fmt.Println(colors)
	fmt.Println(colors[2])
	fmt.Println(len(colors))

}
