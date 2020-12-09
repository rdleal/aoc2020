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

	fmt.Println(FindWeakness(25, list))
}

func FindWeakness(preamble int, list []int) int {
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

	var min, max int

outter:
	for i, num := range list {
		min = num
		max = num
		for _, num2 := range list[i+1:] {
			num += num2

			if num2 > max {
				max = num2
			}

			if num2 < min {
				min = num2
			}

			switch {
			case num > invalidNum:
				break
			case num == invalidNum:
				break outter
			}
		}
	}

	return min + max
}
