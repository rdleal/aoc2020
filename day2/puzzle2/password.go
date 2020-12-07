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

	pos1, pos2, letter := parsePolicy(policy)

	var posCount int
	for i, l := range pass {
		if l == letter && (i == pos1-1 || i == pos2-1) {
			posCount++
		}
	}

	return posCount == 1
}

func splitPolicyAndPass(s string) (policty, pass string) {
	parts := strings.Split(s, ": ")
	return parts[0], parts[1]
}

func parsePolicy(p string) (pos1, pos2 int, letter rune) {
	fields := strings.Fields(p)
	letter = []rune(fields[1])[0]

	parts := strings.Split(fields[0], "-")

	pos1, _ = strconv.Atoi(parts[0])
	pos2, _ = strconv.Atoi(parts[1])

	return
}
