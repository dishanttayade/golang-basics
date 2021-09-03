package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userrating string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name : ")
	name, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate us bet 1 to 5 :")
	userrating, _ = reader.ReadString('\n')
	mynumRate, _ := strconv.ParseFloat(strings.TrimSpace(userrating), 64)

	fmt.Printf("Hello, %v. \nThanks for rating %v stars at %v", name, mynumRate, time.Now().Format(time.Kitchen))

	if mynumRate == 5 {
		fmt.Println("\nYou get a free coupon.")
	}
}
