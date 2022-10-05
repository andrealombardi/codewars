package thirteen

import (
	"strconv"
)

var p = [6]int{1, 10, 9, 12, 3, 4}

func Thirt(n int) int {

	for i := 0; i < 3; i++ {
		n = conv(n)
	}
	return n
}

func conv(n int) (sum int) {

	s := []byte(strconv.Itoa(n))
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		digit, _ := strconv.Atoi(string(s[i]))
		sum += digit * p[j%6]
	}
	return sum
}
