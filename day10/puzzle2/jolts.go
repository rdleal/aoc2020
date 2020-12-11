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

	c := AdaptersPermutations(0, adaptersJolts)

	fmt.Println(c)
}

func AdaptersPermutations(outletJolts int, adaptersJolts []int) int {
	sort.IntSlice(adaptersJolts).Sort()

	cache := make(map[int]int)

	for i, jolts := range adaptersJolts {
		cache[jolts] = i
	}

	permutations := make([]int, len(adaptersJolts))
	permutations[len(adaptersJolts)-1] = 1

	for i := len(adaptersJolts) - 2; i >= 0; i-- {
		sum := 0
		for diff := 1; diff <= 3; diff++ {
			if pos, ok := cache[adaptersJolts[i]+diff]; ok {
				sum += permutations[pos]
			}
		}
		permutations[i] = sum
	}

	var ret int
	for v := 1; v <= 3; v++ {
		if pos, ok := cache[v]; ok {
			ret += permutations[pos]
		}
	}

	return ret
}
