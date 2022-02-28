package isint

import (
	"strconv"
	"testing"
)

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return true
}

func TestIsNumberSuccess(t *testing.T) {
	num := "123"
	isNumber := IsNumber(num)
	if !isNumber {
		t.Errorf("IsNumber(%s) == %t, want %t", num, isNumber, true)
	}
}

func TestIsNumberTableDriven(t *testing.T) {
	var cases = []struct {
		num      string
		expected bool
	}{
		{"123", true},
		{"12b", false},
		{"0", true},
		{"-1", true},
		{"1.1", false},
	}

	for _, c := range cases {
		isNumber := IsNumber(c.num)
		if isNumber != c.expected {
			t.Errorf("IsNumber(%s) == %t, want %t", c.num, isNumber, c.expected)
		}
	}
}
