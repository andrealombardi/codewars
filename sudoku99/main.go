package main

import "fmt"

const total = 45
const length = 9
const cell = 3

func ValidateSolution(m [][]int) bool {

	// validate
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if m[i][j] < 1 || m[i][j] > 9 {
				return false
			}
		}
	}
	// rows
	for i := 0; i < length; i++ {
		sum := 0
		for j := 0; j < length; j++ {
			sum += m[i][j]
			if j == length-1 && sum != total {
				return false
			}
		}
	}

	// columns
	for i := 0; i < length; i++ {
		sum := 0
		for j := 0; j < length; j++ {
			sum += m[j][i]
			if j == length-1 && sum != total {
				return false
			}
		}
	}

	// cells
	for i := 0; i < length; i += 3 {
		for j := 0; j < length; j += 3 {
			sum := 0
			for s := i; s < cell+i; s++ {
				for t := j; t < cell+j; t++ {
					sum += m[s][t]
					if s == cell+j-1 && t == cell+j-1 && sum != 45 {
						return false
					}
				}

			}
		}
	}

	return true
}

func main() {
	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	var testFalse = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 0, 3, 4, 8},
		{1, 0, 0, 3, 4, 2, 5, 6, 0},
		{8, 5, 9, 7, 6, 1, 0, 2, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 0, 1, 5, 3, 7, 2, 1, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 0, 0, 4, 8, 1, 1, 7, 9},
	}

	fmt.Println("vim-go")
	fmt.Println(ValidateSolution(testTrue))
	fmt.Println(ValidateSolution(testFalse))
}
