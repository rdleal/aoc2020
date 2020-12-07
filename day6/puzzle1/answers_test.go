package main

import "testing"

func TestCountYesAnswer(t *testing.T) {
	testCases := []struct {
		name  string
		group []string
		want  int
	}{
		{
			name:  "Group1",
			group: []string{"abc"},
			want:  3,
		},
		{
			name:  "Group2",
			group: []string{"a", "b", "c"},
			want:  3,
		},
		{
			name:  "Group3",
			group: []string{"ab", "ac"},
			want:  3,
		},
		{
			name:  "Group4",
			group: []string{"a", "a", "a", "a"},
			want:  1,
		},
		{
			name:  "Group5",
			group: []string{"b"},
			want:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountYesAnswer(tc.group)

			if got != tc.want {
				t.Errorf("CountYesAnswer(%+q) = got %d; want %d.", tc.group, got, tc.want)
			}
		})
	}
}
