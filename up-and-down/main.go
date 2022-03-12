package main

import (
	"fmt"
	"strings"
)

func Arrange(s string) string {
	l := strings.Split(s, " ")

	for i := 0; i < len(l); i++ {
		if i%2 == 0 {
			if i < len(l)-1 && len(l[i]) > len(l[i+1]) {
				l[i], l[i+1] = l[i+1], l[i]
			}
			l[i] = strings.ToLower(l[i])
		} else {
			if i < len(l)-1 && len(l[i]) < len(l[i+1]) {
				l[i], l[i+1] = l[i+1], l[i]
			}
			l[i] = strings.ToUpper(l[i])
		}
	}

	return fmt.Sprint(strings.Join(l, " "))
}

func main() {
	fmt.Println(Arrange("who hit retaining The That a we taken"))
}
