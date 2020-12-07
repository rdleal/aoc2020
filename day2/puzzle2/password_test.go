package main

import "testing"

func TestValidPassword(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Valid",
			input: "1-3 a: abcde",
			want:  true,
		},
		{
			name:  "Invalid",
			input: "1-3 b: cdefg",
			want:  false,
		},
		{
			name:  "Invalid",
			input: "2-9 c: ccccccccc",
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidPassword(tc.input)

			if got != tc.want {
				t.Errorf("ValidPasswords(%q) = got %t; want %t.", tc.input, got, tc.want)
			}
		})
	}
}
