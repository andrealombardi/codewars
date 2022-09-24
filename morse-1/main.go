package main

import (
	"fmt"
	"strings"
)

var decodeMap = map[string]string{
	".-":   "A",
	"-...": "B",
	"-.-.": "C",
	"-..":  "D",
	".":    "E",
	"..-.": "F",
	"--.":  "G",
	"....": "H",
	"..":   "I",
	".---": "J",
	"-.-":  "K",
	".-..": "L",
	"--":   "M",
	"-.":   "N",
	"---":  "O",
	".--.": "P",
	"--.-": "Q",
	".-.":  "R",
	"...":  "S",
	"-":    "T",
	"..-":  "U",
	"...-": "V",
	".--":  "W",
	"-..-": "X",
	"-.--": "Y",
	"--..": "Z",

	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",

	".-.-.-": ".",
	"--..--": ",",
	"..--..": "?",
	"-...-":  "=",
	"-.-.--": "!",

	"...---...": "SOS",
}

func DecodeMorse(morseCode string) string {
	sentence := make([]string, 0)
	words := strings.Split(strings.Trim(morseCode, " "), "   ")
	for _, word := range words {
		codes := strings.Split(word, " ")
		letters := make([]string, len(codes))
		for i, code := range codes {
			letters[i] = decodeMap[code]
		}
		sentence = append(sentence, strings.Join(letters, ""))

	}
	return strings.Join(sentence, " ")
}

func main() {
	fmt.Println(DecodeMorse(".... . -.-- .--- ..- -.. ."))
}
