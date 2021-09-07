package main

import (
	"fmt"
)

func main() {
	st := 1
	things := []string{"maleta", "bolso", "boleto", "vaso", "casa"}
	fmt.Println(things)

	for i := 0; i < 10; i++ {
		fmt.Println(i + st)
	}
	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	for i := range things {
		fmt.Println(things[i])
	}

	for st < 100 {
		st += st
		// if st==32{
		// 	break
		// }
		// if st == 32 {
		// 	continue
		// }
		if st == 32 {
			goto lcolable
		}
		fmt.Println("Start is now: ", st)
	}

lcolable:
	fmt.Println("http://github.com/dishanttayade")

}
