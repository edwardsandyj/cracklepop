package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 100)

	for i, _ := range arr {
		if (i+1)%3 == 0 {
			if (i+1)%5 == 0 {
				fmt.Println("CracklePop")
			} else {
				fmt.Println("Crackle")
			}
		} else if (i+1)%5 == 0 {
			fmt.Println("Pop")
		} else {
			fmt.Printf("%d\n", i+1)
		}
	}
	return
}
