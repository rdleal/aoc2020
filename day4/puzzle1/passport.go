package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// dropCR drops a terminal \r from the data.
// Code from bufio package.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanPassport(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Scan until single blank line, marking end of passport.
	for width, i := 0, 0; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if d := dropCR(data[:i]); r == '\n' && d[i-1] == '\n' {
			return i + width, dropCR(data[0 : i-1]), nil
		}
	}

	if atEOF {
		return len(data), dropCR(data), nil
	}

	// Request more data.
	return 0, nil, nil
}

func main() {
	var validPassports int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(ScanPassport)
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
