package main

import "testing"

func TestCanContain(t *testing.T) {
	testCases := []struct {
		name  string
		rules []string
		want  int
	}{
		{
			name: "Example",
			rules: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CanContain(tc.rules, "shiny gold")

			if got != tc.want {
				t.Errorf(`CanContain("shiny gold") = got %d; want %d.`, got, tc.want)
			}
		})
	}

}
