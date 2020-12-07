package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var area [][]rune

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		area = append(area, []rune(scanner.Text()))
	}

	fmt.Println(TreesInMap(area))
}

// rune representing a tre in the area.
const tree = '#'

type Map struct {
	area      [][]rune
	bottom    int
	rightEdge int
	posX      int
	posY      int
}

func NewMap(area [][]rune) *Map {
	return &Map{
		area:      area,
		bottom:    len(area) - 1,
		rightEdge: len(area[0]),
	}
}

func (m *Map) Walk(right, down int) (foundTree, bottom bool) {
	m.posX = ((m.posX + right) % m.rightEdge)
	m.posY = m.posY + down

	bottom = m.posY == m.bottom
	foundTree = m.area[m.posY][m.posX] == tree

	return
}

func TreesInMap(area [][]rune) int {
	var (
		bottom bool
		trees  int
	)

	m := NewMap(area)

	for !bottom {
		var foundTree bool
		foundTree, bottom = m.Walk(3, 1)
		if foundTree {
			trees++
		}
	}

	return trees
}
