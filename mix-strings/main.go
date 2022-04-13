package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func Mix(s1, s2 string) string {
	m1 := toMap(s1)
	m2 := toMap(s2)

	allRunes := make(map[rune]bool)
	for k := range m1 {
		allRunes[k] = true
	}
	for k := range m2 {
		allRunes[k] = true
	}

	significantEntries := []entry{}

	for k := range allRunes {
		if m1[k] > 1 || m2[k] > 1 {
			if m1[k] > m2[k] {
				significantEntries = append(significantEntries, entry{string(k), m1[k], "1:"})
			} else if m2[k] > m1[k] {
				significantEntries = append(significantEntries, entry{string(k), m2[k], "2:"})
			} else {
				significantEntries = append(significantEntries, entry{string(k), m2[k], "=:"})
			}
		}
	}

	sort.Slice(significantEntries, func(i, j int) bool {
		if significantEntries[i].Count == significantEntries[j].Count {
			return significantEntries[i].Prefix+significantEntries[i].Value < significantEntries[j].Prefix+significantEntries[j].Value
		}
		return significantEntries[i].Count > significantEntries[j].Count
	})

	entriesAsStrings := []string{}
	for _, v := range significantEntries {
		entriesAsStrings = append(entriesAsStrings, v.String())

	}
	return strings.Join(entriesAsStrings, "/")
}

func toMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		if unicode.IsLower(r) {
			m[r] += 1
		}
	}
	return m
}

type entry struct {
	Value  string
	Count  int
	Prefix string
}

func (e entry) String() string {
	return fmt.Sprintf(e.Prefix + strings.Repeat(e.Value, e.Count))
}

func main() {
	a := Mix("looping is fun but dangerous", "less dangerous than coding")
	fmt.Println(a)
}
