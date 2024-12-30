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
			input:    "hello   world    golang",
			expected: []string{"hello", "world", "golang"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "     ",
			expected: []string{},
		},
		{
			input:    "   hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello   ",
			expected: []string{"hello"},
		},
		{
			input:    "   hello    ",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "hello\t\tworld   golang",
			expected: []string{"hello", "world", "golang"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Fatalf("The length is not same as expected, hence failed!")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Fatalf("The words don't match!, hence failed!")
			}
		}
	}
}
