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

	res := 1

	steps := []struct{ right, down int }{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
	for _, step := range steps {
		res *= TreesInMap(area, step.right, step.down)
	}

	fmt.Println(res)
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

func TreesInMap(area [][]rune, right, down int) int {
	var (
		bottom bool
		trees  int
	)

	m := NewMap(area)

	for !bottom {
		var foundTree bool
		foundTree, bottom = m.Walk(right, down)
		if foundTree {
			trees++
		}
	}

	return trees
}
