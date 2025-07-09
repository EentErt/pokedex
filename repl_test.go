package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "charmander ChArMeLeOn CHARIZARD",
			expected: []string{"charmander", "charmeleon", "charizard"},
		},
		{
			input:    "Ash Ketchum from Pallet Town",
			expected: []string{"ash", "ketchum", "from", "pallet", "town"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of output: %v does not match expected length: %v", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("output word %s does not match expected output %s", word, expectedWord)
				continue
			}
		}
	}
}
