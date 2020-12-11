package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var adaptersJolts []int
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		n, _ := strconv.Atoi(scanner.Text())
		adaptersJolts = append(adaptersJolts, n)

	}

	diffs := CalculateJoltsDiffs(0, adaptersJolts)

	fmt.Println(diffs[1] * diffs[3])
}

func CalculateJoltsDiffs(outletJolts int, adaptersJolts []int) map[int]int {
	diffs := make(map[int]int)

	sort.IntSlice(adaptersJolts).Sort()
	for i := 0; i < len(adaptersJolts); i++ {
		for _, jolts := range adaptersJolts {
			switch diff := jolts - outletJolts; diff {
			case 1, 2, 3:
				outletJolts = jolts
				diffs[diff]++
				break
			}
		}
	}

	// device jolts it allways 3 jolts.
	diffs[3]++

	return diffs
}
