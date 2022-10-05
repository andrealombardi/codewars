package phone

import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	var v []interface{}
	for _, n := range numbers {
		v = append(v, n)
	}
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", v...)
}
