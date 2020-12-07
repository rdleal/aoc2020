package main

import "testing"

func TestTresInMap(t *testing.T) {
	area := [][]rune{
		{'.', '.', '#', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '#', '.', '.', '#', '.'},
		{'.', '.', '#', '.', '#', '.', '.', '.', '#', '.', '#'},
		{'.', '#', '.', '.', '.', '#', '#', '.', '.', '#', '.'},
		{'.', '.', '#', '.', '#', '#', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '#', '.', '#', '.', '.', '.', '.', '#'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '.', '#', '#', '.', '.', '.', '#', '.', '.', '.'},
		{'#', '.', '.', '.', '#', '#', '.', '.', '.', '.', '#'},
		{'.', '#', '.', '.', '#', '.', '.', '.', '#', '.', '#'},
	}
	testCases := []struct {
		name  string
		right int
		down  int
		want  int
	}{
		{
			name:  "1Right2Down",
			right: 1,
			down:  1,
			want:  2,
		},
		{
			name:  "3Right1Down",
			right: 3,
			down:  1,
			want:  7,
		},
		{
			name:  "5Right1Down",
			right: 5,
			down:  1,
			want:  3,
		},
		{
			name:  "7Right1Down",
			right: 7,
			down:  1,
			want:  4,
		},
		{
			name:  "7Right1Down",
			right: 1,
			down:  2,
			want:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := TreesInMap(area, tc.right, tc.down)

			if got != tc.want {
				t.Errorf("TreesInMap(area, %d, %d) = got %d; want %d.", got, tc.right, tc.down, tc.want)
			}
		})
	}
}
