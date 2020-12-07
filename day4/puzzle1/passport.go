package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rdleal/aoc2020/pkg/scan"
)

func main() {
	var validPassports int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(scan.Groups)
	for scanner.Scan() {
		if valid := ValidPassport(scanner.Text()); valid {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID) -- optional
func ValidPassport(p string) bool {
	fields := strings.Fields(p)

	var validFieldsCount int

	for _, field := range fields {
		switch f := strings.Split(field, ":"); f[0] {
		case "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid":
			validFieldsCount++
		}
	}

	return validFieldsCount == 7
}
