package reverse_string

import (
	"testing"
)

func Test1(t *testing.T) {
	input := "Hello world!"
	want := "!dlrow olleH"
	if got := ReverseString(input); got != want {
		t.Errorf("ReverseString() = %s, want %s", got, want)
	}
}
