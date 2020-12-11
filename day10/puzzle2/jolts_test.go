package main

import (
	"reflect"
	"testing"
)

func TestAdaptersPermutations(t *testing.T) {
	testCases := []struct {
		name          string
		outletJolts   int
		adaptersJolts []int
		want          int
	}{
		{
			name:        "Example",
			outletJolts: 0,
			adaptersJolts: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			},
			want: 8,
		},
		{
			name:        "LongerExample",
			outletJolts: 0,
			adaptersJolts: []int{
				28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			want: 19208,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := AdaptersPermutations(tc.outletJolts, tc.adaptersJolts)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("AdaptersPermutations(%d, %v) = got %d; want %d.", tc.outletJolts, tc.adaptersJolts, got, tc.want)
			}
		})
	}
}
