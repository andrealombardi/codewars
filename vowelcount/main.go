package main

import (
	"fmt"
)

var vowels = map[string]int8{
	"a": 0,
	"e": 0,
	"i": 0,
	"o": 0,
	"u": 0,
}

func GetCount(str string) (count int) {
	for _, c := range str {
		_, contains := vowels[string(c)]
		if contains {
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("There are %d vowels\n", GetCount("andrea"))
}



//Best seen on codewars:
func GetCount1(str string) (count int) {
  r := regexp.MustCompile("[aeiou]")
  vowels := r.FindAllString(str, -1)
  return len(vowels)
}

func GetCount2(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}



