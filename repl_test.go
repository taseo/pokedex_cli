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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "       ",
			expected: []string{},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "hello\nworld",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("For input %q expected %d words, but got %d words. Output: %v", c.input, len(c.expected), len(actual), actual)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("For input %q at index %d expected %q, but got %q. Output: %v", c.input, i, expectedWord, word, actual)
			}
		}
	}
}
