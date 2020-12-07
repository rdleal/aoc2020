package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rules []string

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		rules = append(rules, scanner.Text())
	}

	fmt.Println(BagSize(rules, "shiny gold"))
}

func BagSize(rules []string, bag string) int {
	bags := findBags(rules, bag)

	return size(bags, bag)
}

func findBags(rules []string, bag string) map[string]map[string]int {
	bags := make(map[string]map[string]int)

	for _, rule := range rules {
		parts := strings.Split(rule, " contain ")
		curBag := strings.TrimSuffix(parts[0], " bags")

		containBags := strings.Split(parts[1], ", ")
		for _, b := range containBags {
			n, _ := strconv.Atoi(string(b[0]))
			i := strings.Index(b, " bag")

			m, ok := bags[curBag]
			if !ok {
				m = make(map[string]int)
			}
			m[b[2:i]] = n

			bags[curBag] = m
		}
	}
	return bags
}

func size(bags map[string]map[string]int, bag string) int {
	var total int

	for inner, num := range bags[bag] {
		total += num + num*size(bags, inner)
	}

	return total
}
