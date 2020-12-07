package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var rules []string

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		rules = append(rules, scanner.Text())
	}

	fmt.Println(CanContain(rules, "shiny gold"))
}

func CanContain(rules []string, bag string) int {
	bags := make(map[string]struct{})
	findBags(rules, bag, bags)

	return len(bags)
}

func findBags(rules []string, bag string, alreadyFound map[string]struct{}) {
	var canContain []string

	for _, rule := range rules {
		parts := strings.Split(rule, " contain ")
		curBag := strings.TrimSuffix(parts[0], " bags")

		if curBag == bag {
			continue
		}

		if _, ok := alreadyFound[curBag]; !ok && strings.Contains(parts[1], bag) {
			canContain = append(canContain, curBag)
		}
	}

	for _, outterBag := range canContain {
		alreadyFound[outterBag] = struct{}{}
		findBags(rules, outterBag, alreadyFound)
	}
}
