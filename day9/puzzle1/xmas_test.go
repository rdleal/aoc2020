package main

import "testing"

func TestFindInvalidNumber(t *testing.T) {
	testCases := []struct {
		name     string
		preamble int
		list     []int
		want     int
	}{
		{
			name:     "Example",
			preamble: 5,
			list: []int{
				35,
				20,
				15,
				25,
				47,
				40,
				62,
				55,
				65,
				95,
				102,
				117,
				150,
				182,
				127,
				219,
				299,
				277,
				309,
				576,
			},
			want: 127,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindInvalidNumber(tc.preamble, tc.list)

			if got != tc.want {
				t.Errorf("FindNumber(%d, %v) = got %d; want %d.", tc.preamble, tc.list, got, tc.want)
			}
		})
	}
}
