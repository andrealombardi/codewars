package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func NextBigger(n int) int {
	head := strings.Split(fmt.Sprint(n), "")
	tail := make([]string, 0)

	for {

		if len(head) < 2 {
			return -1
		}

		tail = append(tail, head[len(head)-1])
		sort.Strings(tail)
		head = head[:len(head)-1]

		for i := 0; i < len(tail); i++ {
			if tail[i] > head[len(head)-1] {
				tail[i], head[len(head)-1] = head[len(head)-1], tail[i]
				sort.Strings(tail)
				s := strings.Join(head, "") + strings.Join(tail, "")
				result, _ := strconv.Atoi(s)
				return result
			}
		}
	}

}
func main() {
	fmt.Println(NextBigger(12345))
}
