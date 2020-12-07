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
			name:     "Invalid",
			passport: "eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
			want:     false,
		},
		{
			name:     "INvalid",
			passport: "iyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946",
			want:     false,
		},
		{
			name:     "Invalid",
			passport: "hcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
			want:     false,
		},
		{
			name:     "Invalid",
			passport: "hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007",
			want:     false,
		},
		{
			name:     "Valid",
			passport: "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f",
			want:     true,
		},
		{
			name:     "Valid",
			passport: "eyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
			want:     true,
		},
		{
			name:     "Valid",
			passport: "hcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022",
			want:     true,
		},
		{
			name:     "Valid",
			passport: "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			want:     true,
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
