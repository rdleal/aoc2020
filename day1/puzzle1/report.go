package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputs []int

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		txt := scanner.Text()
		n, err := strconv.Atoi(txt)
		if err != nil {
			panic(fmt.Sprintf("input %s is not a number", txt))
		}

		inputs = append(inputs, n)
	}

	fmt.Println(Report(inputs))
}

const delta = 2020

func Report(expenses []int) int {
	diffs := make(map[int]int)

	for _, exp := range expenses {
		if n, ok := diffs[exp]; ok {
			return n * exp
		}
		diffs[delta-exp] = exp
	}

	return 0
}
