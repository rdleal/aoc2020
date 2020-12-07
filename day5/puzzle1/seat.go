package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var highestSeatID int

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		boardingPass := scanner.Text()

		if _, _, seatID := Seat(boardingPass); seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Println(highestSeatID)
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
