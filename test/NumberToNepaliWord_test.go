package test

import (
	"strconv"
	"testing"

	nepali_utils "github.com/PrameshKarki/nepali_utils/lib"
)

func TestNumberToNepaliWord(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{12345, "बाह्र हजार तीन सय पैँतालीस"},
		{500, "पाँच सय"},
		{100, "एक सय"},
		{12312, "बाह्र हजार तीन सय बाह्र"},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {
			result := nepali_utils.NumberToNepaliWord(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
