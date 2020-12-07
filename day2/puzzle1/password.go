package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var validCount int

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		if ok := ValidPassword(scanner.Text()); ok {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func ValidPassword(in string) bool {
	policy, pass := splitPolicyAndPass(in)

	min, max, letter := parsePolicy(policy)

	var letterCount int
	for _, l := range pass {
		if l == letter {
			letterCount++
		}
	}

	return letterCount >= min && letterCount <= max
}

func splitPolicyAndPass(s string) (policty, pass string) {
	parts := strings.Split(s, ":")
	return parts[0], parts[1]
}

func parsePolicy(p string) (min, max int, letter rune) {
	fields := strings.Fields(p)
	letter = []rune(fields[1])[0]

	parts := strings.Split(fields[0], "-")

	min, _ = strconv.Atoi(parts[0])
	max, _ = strconv.Atoi(parts[1])

	return
}
