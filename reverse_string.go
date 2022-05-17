package reverse_string

import (
	"strings"
)

func ReverseString(input string) (output string) {
	for _, s := range strings.Split(input, "") {
		output = s + output
	}
	return output
}
