package main

import "testing"

func TestScannerScan(t *testing.T) {
	testCases := []struct {
		name     string
		src      string
		expected []string
	}{
		{
			name:     "scan single chars",
			src:      "(){},.-+;/*",
			expected: []string{"(", ")", "{", "}", ",", ".", "-", "+", ";", "/", "*"},
		},
	}

	for _, tc := range testCases {
		scanner := newScanner(tc.src)
		if err := scanner.scan(); err != nil {
			t.Fatal(err)
		}

		if len(scanner.tokens) != len(tc.expected) {
			t.Fatalf("%s: expected %d tokens, got %d", tc.name, len(tc.expected), len(scanner.tokens))
		}

		for i, text := range tc.expected {
			if scanner.tokens[i].Text != text {
				t.Fatalf("%s: expected %s at %d, got %s", tc.name, text, i, scanner.tokens[i].Text)
			}
		}
	}
}
