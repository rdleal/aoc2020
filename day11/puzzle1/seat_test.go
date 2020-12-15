package main

import "testing"

func TestOccupiedSeats(t *testing.T) {
	testCases := []struct {
		name  string
		seats [][]rune
		want  int
	}{
		{
			name: "Example",
			seats: [][]rune{
				{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
				{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
				{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
				{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
			},
			want: 37,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := OccupiedSeats(tc.seats)

			if got != tc.want {
				t.Errorf("OccupiedSeats() = got %d; want %d.", got, tc.want)
			}
		})
	}
}
