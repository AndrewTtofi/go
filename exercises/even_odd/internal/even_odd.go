package internal

import (
	"fmt"
	"strconv"
)

func EvenOdd(size int) []string {
	var numRange []int
	var results []string
	for i := 0; i < size; i++ {
		numRange = append(numRange, i)
	}
	for i := range numRange {
		if numRange[i]%2 == 0 {
			results = append(results, fmt.Sprintf("The number %s is even\n", strconv.Itoa(numRange[i])))
		} else {
			results = append(results, fmt.Sprintf("The number %s is odd\n", strconv.Itoa(numRange[i])))
		}
	}
	return results
}

func SpecificNum(num int) string {
	var result string
	if num%2 == 0 {
		result = fmt.Sprintf("The number %d is even", num)
	} else {
		result = fmt.Sprintf("The number %d is odd", num)
	}
	return result
}
