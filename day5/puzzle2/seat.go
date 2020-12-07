package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var seats []int

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		boardingPass := scanner.Text()

		_, _, seatID := Seat(boardingPass)
		seats = append(seats, seatID)
	}

	sorted := sort.IntSlice(seats)
	sorted.Sort()

	var prevSeat int
	for _, seat := range sorted {
		if prevSeat > 0 && prevSeat+1 != seat {
			fmt.Printf("Missing seat: %d\n", prevSeat+1)
		}
		prevSeat = seat
	}
}

func Seat(boardingPass string) (row, column, id int) {
	row = position(boardingPass[0:7], 0, 127, 'F', 'B')
	column = position(boardingPass[7:], 0, 7, 'L', 'R')

	id = row*8 + column

	return
}

func position(s string, lower, upper int, front, back rune) int {
	var pos int

	for _, half := range s {
		switch half {
		case front:
			upper -= (upper - lower + 1) / 2
			pos = lower
		case back:
			lower += (upper + 1 - lower) / 2
			pos = upper
		}
	}

	return pos
}
