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
		diffs[exp] = delta - exp
		for n, diff := range diffs {
			n2 := diff - exp
			if _, ok := diffs[n2]; ok {
				return exp * n * n2
			}
		}
	}

	return 0
}
