package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	l := strings.Split(str, " ")
	for i := range l {
		if len(l[i]) > 4 {
			l[i] = reverse(l[i])
		}
	}

	return strings.Join(l, " ")
} // SpinWords

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(SpinWords("Welcome"))
}
