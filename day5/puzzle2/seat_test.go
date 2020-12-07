package main

import "testing"

func TestSeat(t *testing.T) {
	testCases := []struct {
		name         string
		boardingPass string
		wantRow      int
		wantColumn   int
		wantID       int
	}{
		{
			boardingPass: "FBFBBFFRLR",
			wantRow:      44,
			wantColumn:   5,
			wantID:       357,
		},
		{
			boardingPass: "BFFFBBFRRR",
			wantRow:      70,
			wantColumn:   7,
			wantID:       567,
		},
		{
			boardingPass: "FFFBBBFRRR",
			wantRow:      14,
			wantColumn:   7,
			wantID:       119,
		},
		{
			boardingPass: "BBFFBBFRLL",
			wantRow:      102,
			wantColumn:   4,
			wantID:       820,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.boardingPass, func(t *testing.T) {
			gotRow, gotColumn, gotID := Seat(tc.boardingPass)

			if gotRow != tc.wantRow {
				t.Errorf("Seat(%q) = got row: %d; want %d.", tc.boardingPass, gotRow, tc.wantRow)
			}

			if gotColumn != tc.wantColumn {
				t.Errorf("Seat(%q) = got column: %d; want %d.", tc.boardingPass, gotColumn, tc.wantColumn)
			}

			if gotID != tc.wantID {
				t.Errorf("Seat(%q) = got ID: %d; want %d.", tc.boardingPass, gotID, tc.wantID)
			}
		})
	}
}
