package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

type Coord = [2]int

var offsets = [][]Coord{
	{{0, 1}, {0, 2}, {0, 3}},
	{{0, -1}, {0, -2}, {0, -3}},
	{{1, 1}, {2, 2}, {3, 3}},
	{{1, -1}, {2, -2}, {3, -3}},
	{{-1, -1}, {-2, -2}, {-3, -3}},
	{{-1, 1}, {-2, 2}, {-3, 3}},
	{{-1, 0}, {-2, 0}, {-3, 0}},
	{{1, 0}, {2, 0}, {3, 0}},
}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	chart := make([][]string, 0)

	for _, line := range lines {
		chart = append(chart, strings.Split(line, ""))
	}

	xmasCount := 0
	for y, line := range chart {
		for x, char := range line {
			if char != "X" {
				continue
			}

			for _, word := range getWords(chart, y, x) {
				if word == "XMAS" {
					xmasCount++
				}
			}
		}
	}

	fmt.Printf("XMAS: %d\n", xmasCount)
}

func getWords(chart [][]string, startY, startX int) []string {
	height := len(chart)
	width := len(chart[0])

	words := make([]string, 0, len(offsets))

offsets:
	for _, offset := range offsets {
		var word strings.Builder
		word.WriteString(chart[startY][startX])

		for _, coord := range offset {
			y := startY - coord[0]
			x := startX - coord[1]

			if y >= height || y < 0 || x >= width || x < 0 {
				continue offsets
			}

			word.WriteString(chart[y][x])
		}

		words = append(words, word.String())
	}

	return words
}
