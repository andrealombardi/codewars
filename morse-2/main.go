package main

import (
	"fmt"
	"strings"
)

func DecodeBits(bits string) string {

	bits = strings.Trim(bits, "0")
	unit := detectUnit(bits)

	morseCode := make([]string, 0)
	words := strings.Split(bits, strings.Repeat("0", 7*unit))
	for _, word := range words {
		morseWord := make([]string, 0)
		characters := strings.Split(word, strings.Repeat("0", 3*unit))
		for _, character := range characters {
			signs := strings.Split(character, strings.Repeat("0", unit))
			morse := make([]string, len(signs))

			for i, sign := range signs {
				switch len(sign) {
				case 1 * unit:
					morse[i] = "."
				case 3 * unit:
					morse[i] = "-"
				default:
					fmt.Printf("unexpected length: %d", len(sign))
				}
			}
			morseWord = append(morseWord, strings.Join(morse, ""))
		}
		morseCode = append(morseCode, strings.Join(morseWord, " "))

	}
	return strings.Join(morseCode, "   ")
}

func DecodeMorse(morseCode string) string {
	sentence := make([]string, 0)
	words := strings.Split(strings.Trim(morseCode, " "), "   ")
	for _, word := range words {
		codes := strings.Split(word, " ")
		letters := make([]string, len(codes))
		for i, code := range codes {
			letters[i] = MORSE_CODE[code]
		}
		sentence = append(sentence, strings.Join(letters, ""))

	}
	return strings.Join(sentence, " ")
}

func detectUnit(bits string) int {
	unit := len(bits)
	count := 0
	countZero := true
	if strings.Contains(bits, "0") && strings.Contains(bits, "1") {

		for _, r := range bits {

			if r == '0' {
				if countZero {
					count++
				} else {
					if count > 0 && count < unit {
						unit = count
					}
					countZero = true
					count = 1
				}
			} else if r == '1' {
				if !countZero {
					count++
				} else {
					if count > 0 && count < unit {
						unit = count
					}
					countZero = false
					count = 1
				}
			} else {
				panic(r)
			}
		}
	}
	return unit
}

var MORSE_CODE = map[string]string{
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

func main() {
	//fmt.Println(DecodeMorse(DecodeBits("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011")))
	fmt.Println(DecodeMorse(DecodeBits("111")))
}
