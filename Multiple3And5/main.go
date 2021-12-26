package main

import (
	"fmt"
)

func Multiple3And5(number int) int {
	if number < 0 {
		return 0
	}

	result := 0
	for i := 3; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = result + i
		}

	}
	return result
}

//From Gauss sum all the numbers up to 100 -> n(n+1)/2
func BestAnswer(number int) int {
	number--
	sum_3 := (number / 3) * (number/3 + 1) * 3 / 2
	sum_5 := (number / 5) * (number/5 + 1) * 5 / 2
	sum_15 := (number / 15) * (number/15 + 1) * 15 / 2
	return sum_3 + sum_5 - sum_15

}
func main() {
	fmt.Printf("The sum is %d\n", Multiple3And5(10))
}
