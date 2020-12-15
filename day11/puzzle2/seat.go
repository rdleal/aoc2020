package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var seats [][]rune

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		seats = append(seats, []rune(scanner.Text()))
	}

	fmt.Println(OccupiedSeats(seats))
}

func OccupiedSeats(seats [][]rune) int {

	for {
		seen := make(map[int]map[int]rune)

		var prev, cur string
		var occupied int

		for i, row := range seats {
			seen[i] = make(map[int]rune)
			for j, seat := range row {
				seen[i][j] = seat
				prev = prev + string(seat)
			}
		}

		for i, row := range seats {
			for j, seat := range row {
				occupieds := occupiedAdjacents(seen, i, j)

				switch seat {
				case 'L':
					if occupieds == 0 {
						seat = '#'
					}
				case '#':
					if occupieds >= 5 {
						seat = 'L'
					}
				}

				if isOccupied(seat) {
					occupied++
				}

				seats[i][j] = seat
				cur = cur + string(seat)
			}
		}

		if prev == cur {
			return occupied
		}
	}

}

func occupiedAdjacents(seats map[int]map[int]rune, row, col int) int {
	var occupieds int

	up := row - 1
	down := row + 1

	right := col + 1
	left := col - 1

	i := down
	for ; i < len(seats) && !isSeat(seats[i][col]); i++ {
	}

	if isOccupied(seats[i][col]) {
		occupieds++
	}

	i = up
	for ; i > 0 && !isSeat(seats[i][col]); i-- {
	}

	if isOccupied(seats[i][col]) {
		occupieds++
	}

	r := right
	for ; r < len(seats[0]) && !isSeat(seats[row][r]); r++ {
	}

	if isOccupied(seats[row][r]) {
		occupieds++
	}

	l := left
	for ; l > 0 && !isSeat(seats[row][l]); l-- {
	}

	if isOccupied(seats[row][l]) {
		occupieds++
	}

	for i, r = up, right; i > 0 && r < len(seats[0]) && !isSeat(seats[i][r]); i, r = i-1, r+1 {
	}

	if isOccupied(seats[i][r]) {
		occupieds++
	}

	for i, l = up, left; i > 0 && l > 0 && !isSeat(seats[i][l]); i, l = i-1, l-1 {
	}

	if isOccupied(seats[i][l]) {
		occupieds++
	}

	for i, r = down, right; i < len(seats) && r < len(seats[0]) && !isSeat(seats[i][r]); i, r = i+1, r+1 {
	}

	if isOccupied(seats[i][r]) {
		occupieds++
	}

	for i, l = down, left; i < len(seats) && l > 0 && !isSeat(seats[i][l]); i, l = i+1, l-1 {
	}

	if isOccupied(seats[i][l]) {
		occupieds++
	}

	return occupieds

}

func isSeat(r rune) bool {
	return r != '.'
}

func isOccupied(r rune) bool {
	return r == '#'
}
