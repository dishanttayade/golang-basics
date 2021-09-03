package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var Str string
	// fmt.Scanln(&Str)
	// fmt.Println(Str)
	// same problem as C

	var name string = "MyName"
	var a, b = fmt.Println(name)
	fmt.Println(a, b)
	var c, _ = fmt.Println(name)
	fmt.Println(c)

	read := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name: ")
	avar, _ := read.ReadString('\n')
	fmt.Println(avar)

	read1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating : ")
	rating, _ := read1.ReadString('\n')
	myrating, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	fmt.Println(myrating + 2)

}
