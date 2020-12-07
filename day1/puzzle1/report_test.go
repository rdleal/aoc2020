package main

import "testing"

func TestReport(t *testing.T) {
	testCases := []struct {
		name     string
		expenses []int
		want     int
	}{
		{
			name: "Example",
			expenses: []int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			want: 514579,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Report(tc.expenses)

			if got != tc.want {
				t.Errorf("Report() = got %d; want %d.", got, tc.want)
			}
		})
	}
}
