package main

import (
	"fmt"
)

func main() {
	var n []int
	for i := 0; i < 11; i++ {
		n = append(n, i)
	}
	for i := range n {
		if n[i]%2 == 0 {
			fmt.Println("The number ", n[i], " is even")
		} else {
			fmt.Println("The number ", n[i], " is odd")
		}
	}
}
