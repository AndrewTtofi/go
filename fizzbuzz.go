package main

import (
	"fmt"
	"strconv"
)

type fizzbuzz []string

func initCount(max_num int) fizzbuzz {
	f := fizzbuzz{}
	for i := 0; i <= max_num; i++ {
		if i%3 == 0 {
			f = append(f, "Fizz")
		} else if i%5 == 0 {
			f = append(f, "Buzz")
		} else {
			output := strconv.Itoa(i)
			f = append(f, output)
		}
	}
	return f
}

func (f fizzbuzz) print() {
	for i, num := range f {
		fmt.Println(i, num)
	}
}
