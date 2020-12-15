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
					if occupieds >= 4 {
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

	if isOccupied(seats[down][col]) {
		occupieds++
	}

	if isOccupied(seats[up][col]) {
		occupieds++
	}

	if isOccupied(seats[row][right]) {
		occupieds++
	}

	if isOccupied(seats[row][left]) {
		occupieds++
	}

	if isOccupied(seats[up][right]) {
		occupieds++
	}

	if isOccupied(seats[up][left]) {
		occupieds++
	}

	if isOccupied(seats[down][right]) {
		occupieds++
	}

	if isOccupied(seats[down][left]) {
		occupieds++
	}

	return occupieds

}

func isOccupied(r rune) bool {
	return r == '#'
}
