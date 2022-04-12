package main

import (
	"fmt"
)

func Snail(snailMap [][]int) []int {

	if len(snailMap) == 0 || len(snailMap[0]) == 0 {
		return snail
	}

	var snail = make([]int, len(snailMap)*len(snailMap))
	var minX, minY, position int
	var maxX, maxY = len(snailMap), len(snailMap)

	right := func() {
		for i := minX; i < maxX; i++ {
			snail[position] = snailMap[minX][i]
			position++
		}
		minY++
	}

	down := func() {
		maxX--
		for i := minY; i < maxY; i++ {
			snail[position] = snailMap[i][maxX]
			position++
		}
	}

	left := func() {
		maxY--
		for i := maxX - 1; i >= minX; i-- {
			snail[position] = snailMap[maxY][i]
			position++
		}
	}

	up := func() {
		for i := maxY - 1; i >= minY; i-- {
			snail[position] = snailMap[i][minX]
			position++
		}
		minX++
	}

	finished := func() bool {
		return position == len(snail)
	}

	for {
		right()
		if finished() {
			return snail
		}
		down()
		if finished() {
			return snail
		}
		left()
		if finished() {
			return snail
		}
		up()
		if finished() {
			return snail
		}
	}
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

	snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	snailMap1 := [][]int{{1}}

	fmt.Println("vim-go")
	fmt.Println(Snail(testTrue))
	fmt.Println(Snail(testFalse))
	fmt.Println(Snail(snailMap))
	fmt.Println(Snail(snailMap1))
}
