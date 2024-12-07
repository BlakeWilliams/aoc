package main

import (
	_ "embed"
	"errors"
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
	obstacles []Coord
}

var ErrLoop = errors.New("hell yeah")

func (m *Map) Walk() error {
	for {
		next, obstacle, inBounds := m.NextStep()
		if !inBounds {
			return nil
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

			m.obstacles = append(m.obstacles, next)

			repeatSize := 0
			for i := 3; i < len(m.obstacles); i++ {
				comp := m.obstacles[len(m.obstacles)-i]
				if comp[0] == next[0] && comp[1] == next[1] {
					repeatSize = i - 1
					break
				}
			}

			if repeatSize == 0 {
				continue
			}

			if len(m.obstacles) < repeatSize*12 {
				continue
			}

			loop := true
			for i := 0; i < repeatSize; i++ {
				end := len(m.obstacles) - 1
				if m.obstacles[end-i] != m.obstacles[end-i-repeatSize] {
					loop = false
					break
				}
			}

			if loop {
				return ErrLoop
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

	loops := 0
	origGuard := Coord{m.guard[0], m.guard[1]}
	iter := 0
	for y := 0; y < len(m.grid)-1; y++ {
		for x := 0; x <= len(m.grid[0])-1; x++ {
			iter++
			m.obstacles = []Coord{}
			m.visited = map[Coord]struct{}{m.guard: {}}
			m.guard[0], m.guard[1] = origGuard[0], origGuard[1]
			m.direction = DirUp

			origItem := m.grid[y][x]
			m.grid[y][x] = '#'

			if err := m.Walk(); err == ErrLoop {
				loops++
			}

			m.grid[y][x] = origItem
		}
	}

	fmt.Println(loops)
}
