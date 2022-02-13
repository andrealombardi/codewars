package main

import "fmt"

func main() {
	fmt.Println(RemoveChar("vim-go"))
}

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
