package main

import "testing"

func TestBagSize(t *testing.T) {
	testCases := []struct {
		name  string
		rules []string
		want  int
	}{
		{
			name: "Example",
			rules: []string{
				"shiny gold bags contain 2 dark red bags.",
				"dark red bags contain 2 dark orange bags.",
				"dark orange bags contain 2 dark yellow bags.",
				"dark yellow bags contain 2 dark green bags.",
				"dark green bags contain 2 dark blue bags.",
				"dark blue bags contain 2 dark violet bags.",
				"dark violet bags contain no other bags.",
			},
			want: 126,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BagSize(tc.rules, "shiny gold")

			if got != tc.want {
				t.Errorf(`CanContain("shiny gold") = got %d; want %d.`, got, tc.want)
			}
		})
	}

}
