package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var list []int
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		n, _ := strconv.Atoi(scanner.Text())
		list = append(list, n)

	}

	fmt.Println(FindInvalidNumber(25, list))
}

func FindInvalidNumber(preamble int, list []int) int {
	var invalidNum int

	for found, counter := true, 0; found; counter++ {
		preambleMap := make(map[int]struct{})

		preambleList := list[counter : preamble+counter]

		for _, n := range preambleList {
			preambleMap[n] = struct{}{}
		}

		next := list[preamble+counter]

		for _, n := range preambleList {
			var diff int

			switch {
			case n > next:
				diff = n - next
			default:
				diff = next - n
			}

			if diff == n {
				continue
			}

			if _, found = preambleMap[diff]; found {
				break
			}

			invalidNum = next
		}
	}

	return invalidNum
}
