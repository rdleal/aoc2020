package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

type ValidFieldFunc func(v string) bool

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
var fieldValidator = map[string]ValidFieldFunc{
	"byr": func(v string) bool {
		yr, _ := strconv.Atoi(v)

		return yr >= 1920 && yr <= 2002
	},
	"ecl": func(v string) bool {
		switch v {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		default:
			return false
		}
	},
	"eyr": func(v string) bool {
		yr, _ := strconv.Atoi(v)

		return yr >= 2020 && yr <= 2030
	},
	"iyr": func(v string) bool {
		yr, _ := strconv.Atoi(v)
		return yr >= 2010 && yr <= 2020
	},
	"hgt": func(v string) bool {
		s := strings.TrimRight(v, "cmin")
		n, _ := strconv.Atoi(s)

		switch {
		case strings.HasSuffix(v, "cm"):
			return n >= 150 && n <= 193
		case strings.HasSuffix(v, "in"):
			return n >= 59 && n <= 76
		default:
			return false
		}
	},
	"hcl": func(v string) bool {
		if !strings.HasPrefix(v, "#") {
			return false
		}
		h := strings.Map(mapHex, v)

		return len(h) == 6
	},
	"pid": func(v string) bool {
		if len(v) != 9 {
			return false
		}
		_, err := strconv.Atoi(v)
		return err == nil
	},
}

func mapHex(r rune) rune {
	if unicode.IsDigit(r) || r >= 'a' && r <= 'f' {
		return r
	}

	return -1
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
		f := strings.Split(field, ":")
		if validField, ok := fieldValidator[f[0]]; ok && validField(f[1]) {
			validFieldsCount++
		}
	}

	return validFieldsCount == 7
}
