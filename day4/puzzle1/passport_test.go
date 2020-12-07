package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestScanPassport(t *testing.T) {
	testCases := []struct {
		name          string
		batch         string
		wantPassports int
	}{
		{
			name:          "OnePassport",
			batch:         "pid:883116919 iyr:2018 ecl:brn byr:1938 hgt:187cm eyr:2020",
			wantPassports: 1,
		},
		{
			name:          "TwoPassport",
			batch:         "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm\n\niyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929",
			wantPassports: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.batch)
			s := bufio.NewScanner(r)
			s.Split(ScanPassport)

			var gotPassports int
			for ; s.Scan(); gotPassports++ {
				//t.Logf("%q\n", s.Text())
			}

			if gotPassports != tc.wantPassports {
				t.Errorf("Scan() = got %d passports; want %d.", gotPassports, tc.wantPassports)
			}
		})
	}
}

func TestValidPassport(t *testing.T) {
	testCases := []struct {
		name     string
		passport string
		want     bool
	}{
		{
			name:     "AllFields",
			passport: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm",
			want:     true,
		},
		{
			name:     "MissingHeight",
			passport: "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929",
			want:     false,
		},
		{
			name:     "MissingCid",
			passport: "hcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm",
			want:     true,
		},
		{
			name:     "MissingByr",
			passport: "hcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in",
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidPassport(tc.passport)

			if got != tc.want {
				t.Errorf("ValidPassport(%q) = got %t; want %t.", tc.passport, got, tc.want)
			}
		})
	}
}
