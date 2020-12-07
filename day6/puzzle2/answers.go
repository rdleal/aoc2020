package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rdleal/aoc2020/pkg/scan"
)

func main() {
	var totalYes int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(scan.Groups)
	for scanner.Scan() {
		answers := scanner.Text()
		totalYes += CountYesAnswer(strings.Fields(answers))
	}

	fmt.Println(totalYes)
}

func CountYesAnswer(group []string) int {
	uniqAnswers := make(map[rune]int)

	for _, answers := range group {
		for _, answer := range answers {
			uniqAnswers[answer]++
		}
	}

	var total int
	for _, num := range uniqAnswers {
		if num == len(group) {
			total++
		}
	}

	return total
}
