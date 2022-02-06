package kata

import (
	"strconv"
	"strings"
)

func Points(games []string) int {
	points := 0
	for _, game := range games {
		result := strings.Split(game, ":")
		x, _ := strconv.Atoi(result[0])
		y, _ := strconv.Atoi(result[1])

		if x > y {
			points += 3
		} else if x == y {
			points += 1
		}
	}

	return points
}

//Loved this one
func Points2(games []string) (res int) {
	for _, v := range games {
		if v[0] > v[2] {
			res += 3
		}
		if v[0] == v[2] {
			res += 1
		}
	}
	return
}
