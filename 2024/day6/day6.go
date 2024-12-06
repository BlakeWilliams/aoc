package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type Coord = [2]int

type Direction Coord

var (
	DirUp    Direction = Coord{-1, 0}
	DirDown  Direction = Coord{1, 0}
	DirLeft  Direction = Coord{0, -1}
	DirRight Direction = Coord{0, 1}
)

type Map struct {
	grid      [][]rune
	visited   map[Coord]struct{}
	guard     Coord
	direction Direction
}

func (m *Map) Walk() {
	for {
		next, obstacle, inBounds := m.NextStep()
		if !inBounds {
			return
		}

		switch obstacle {
		case '.':
		case '#':
			switch m.direction {
			case DirUp:
				m.direction = DirRight
			case DirRight:
				m.direction = DirDown
			case DirDown:
				m.direction = DirLeft
			case DirLeft:
				m.direction = DirUp
			}

			continue
		}
		m.guard = next
		m.visited[next] = struct{}{}
	}
}

func (m *Map) NextStep() (Coord, rune, bool) {
	next := Coord{m.guard[0] + m.direction[0], m.guard[1] + m.direction[1]}
	if next[0] < 0 || next[0] >= len(m.grid) {
		return next, ' ', false
	}

	if next[1] < 0 || next[1] >= len(m.grid[next[0]]) {
		return next, ' ', false
	}

	return next, m.grid[next[0]][next[1]], true
}

func main() {
	rows := strings.Split(input, "\n")
	m := Map{
		grid:      make([][]rune, 0, len(rows)),
		visited:   make(map[Coord]struct{}),
		direction: DirUp,
	}
	for y, row := range rows {
		positions := []rune(row)
		for x, pos := range positions {
			if pos == '^' {
				m.guard = Coord{y, x}
				m.visited[m.guard] = struct{}{}
			}
		}
		m.grid = append(m.grid, positions)
	}

	m.Walk()

	fmt.Println(len(m.visited))
}
