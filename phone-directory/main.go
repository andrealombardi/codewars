package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Phone(dir, num string) string {

	matchCount := strings.Count(dir, "+"+num)
	if matchCount == 0 {
		return fmt.Sprintf("Error => Not found: %s", num)
	} else if matchCount > 1 {
		return fmt.Sprintf("Error => Too many people: %s", num)
	}

	phoneRegex := regexp.MustCompile(`([^\s]*\+(\d{1,2}-\d{3}-\d{3}-\d{4})[^\s]*)`)
	nameRegex := regexp.MustCompile(`([^\s]*<(.*)>[^\s]*)`)
	multispaceRegex := regexp.MustCompile(`\s{1,}`)
	noSpecialCharactersRegex := regexp.MustCompile(`[^A-Za-z0-9\s\.-]+`)

	lines := strings.Split(dir, "\n")

	for _, line := range lines {
		if strings.Contains(line, "+"+num) {
			phone := phoneRegex.FindAllStringSubmatch(line, -1)
			name := nameRegex.FindAllStringSubmatch(line, -1)
			nameless := strings.ReplaceAll(line, name[0][1], "")
			phoneless := strings.ReplaceAll(nameless, phone[0][1], "")
			specialCharless := noSpecialCharactersRegex.ReplaceAllString(phoneless, " ")
			address := strings.TrimSpace(multispaceRegex.ReplaceAllString(specialCharless, " "))

			return fmt.Sprintf("Phone => %s, Name => %s, Address => %s", phone[0][2], name[0][2], address)

		}

	}
	return ""
}

const ddr = "/+1-541-754-3010 156 Alphand_St. i<J Steeve>i\n 133, Green, Rd. <E Kustur> NY-56423 ;+1-541-914-3010\n" + "+1-541-984-3012 <P Reed> /PO Box 530; Pollocksville, NC-28573\n :+1-321-512-2222 <Paul Dive> Sequoia Alley PQ-67209\n" + "+1-741-984-3090 <Peter Reedgrave> _Chicago\n :+1-921-333-2222 <Anna Stevens> Haramburu_Street AA-67209\n" + "+1-111-544-8973 <Peter Pan> LA\n +1-921-512-2222 <Wilfrid Stevens> Wild Street AA-67209\n" + "<Peter Gone> LA ?+1-121-544-8974 \n <R Steell> Quora Street AB-47209 +1-481-512-2222\n" + "<Arthur Clarke> San Antonio $+1-121-504-8974 TT-45120\n <Ray Chandler> Teliman Pk. !+1-681-512-2222! AB-47209,\n" + "<Sophia Loren> +1-421-674-8974 Bern TP-46017\n <Peter O'Brien> High Street +1-908-512-2222; CC-47209\n" + "<Anastasia> +48-421-674-8974 Via Quirinal    Roma\n <P Salinger> Main Street, +1-098-512-2222, Denver\n" + "<C Powel> *+19-421-674-8974 Chateau des Fosses Strasbourg F-68000\n <Bernard Deltheil> +1-498-512-2222; Mount Av.  Eldorado\n" + "+1-099-500-8000 <Peter Crush> Labrador Bd.\n +1-931-512-4855 <William Saurin> Bison Street CQ-23071\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n"

func main() {
	fmt.Println(Phone(ddr, "48-421-674-8974"))
}