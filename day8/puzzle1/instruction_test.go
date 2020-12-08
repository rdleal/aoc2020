package main

import "testing"

func TestRunProgram(t *testing.T) {
	testCases := []struct {
		name         string
		instructions []string
		want         int
	}{
		{
			name: "Example",
			instructions: []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			},
			want: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := RunProgram(tc.instructions)

			if got != tc.want {
				t.Errorf("RunProgram(%+q) = got accumulator %d; want %d.", tc.instructions, got, tc.want)
			}
		})
	}
}
