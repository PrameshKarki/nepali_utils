package test

import (
	"testing"

	nepali_utils "github.com/PrameshKarki/nepali_utils/lib"
)

func TestEnglishToNepali(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"12345", "१२३४५"},
		{"mero nepal", "मेरो नेपाल"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := nepali_utils.EnglishToNepali(tc.input)
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
