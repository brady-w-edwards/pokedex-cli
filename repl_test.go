package main

import "testing"

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
			input:    "Testing Caps",
			expected: []string{"testing", "caps"},
		},
		{
			input:    "WHAT ABOUT ALL CAPS",
			expected: []string{"what", "about", "all", "caps"},
		},
		{
			input:    " extra  white    space",
			expected: []string{"extra", "white", "space"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Actual length does not equal expected length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("words are not equal \nactual: %s\nexpected: %s", word, expectedWord)
			}
		}
	}
}
