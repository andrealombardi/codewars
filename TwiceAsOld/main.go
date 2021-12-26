package main

import(
	"fmt"
)


func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
  return abs(sonYearsOld - (dadYearsOld - sonYearsOld))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	fmt.Printf("Twice as old: %d\n", TwiceAsOld(42,21))
}
