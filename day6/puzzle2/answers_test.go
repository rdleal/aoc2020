package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestScanPassport(t *testing.T) {
	testCases := []struct {
		name        string
		batch       string
		wantAnswers int
	}{
		{
			name:        "OneAnswer",
			batch:       "a",
			wantAnswers: 1,
		},
		{
			name:        "ThreeAnswers",
			batch:       "abc\n\na\nb\nc\n\nab\nnac",
			wantAnswers: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.batch)
			s := bufio.NewScanner(r)
			s.Split(ScanAnswers)

			var gotAnswers int
			for ; s.Scan(); gotAnswers++ {
				//t.Logf("%q\n", s.Text())
			}

			if gotAnswers != tc.wantAnswers {
				t.Errorf("Scan() = got %d answers; want %d.", gotAnswers, tc.wantAnswers)
			}
		})
	}
}

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
			want:  0,
		},
		{
			name:  "Group3",
			group: []string{"ab", "ac"},
			want:  1,
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
