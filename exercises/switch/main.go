package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Switch for ", i, " as ")
	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}
	//Print if it is weekend or print the weekday
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend time.It's ", time.Now().Weekday())
	default:
		fmt.Println("The day is ", time.Now().Weekday())
	}

	//Print if the time is before lunch or night
	switch {
	case time.Now().Hour() > 12:
		fmt.Println("the time is pm: ", time.Now().Hour())

	default:
		fmt.Println("the time is am: ", time.Now().Hour())
	}

	//create a function and define the type of the variable inserted
	typeoffunc := func(in interface{}) {
		switch x := in.(type) {
		case bool:
			fmt.Println("boolean", x)
		case int:
			fmt.Println("integer", x)
		case string:
			fmt.Println("string", x)
		}
	}
	typeoffunc(true)
	typeoffunc(1)
	typeoffunc("yiannoui")
}
