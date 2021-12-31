package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Interval: %v", IntervalInsert([][2]int{{1, 2}, {3, 4}, {5, 6}}, [2]int{2, 3}))
}

//hinted!
func IntervalInsert(r [][2]int, z [2]int) [][2]int {

	result := make([][2]int, 0)
	for i, interval := range r {
		if interval[1] < z[0] {
			result = append(result, interval)
		} else if interval[0] > z[1] {
			result = append(result, z)
			result = append(result, r[i:]...)
			return result
		} else {
			if interval[0] < z[0] {
				z[0] = interval[0]
			}
			if interval[1] > z[1] {
				z[1] = interval[1]
			}
		}
	}
	return append(result, z)
}
