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
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	{{0, 0}, {0, -1}, {0, -2}, {0, -3}},
	{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
	{{0, 0}, {1, -1}, {2, -2}, {3, -3}},
	{{0, 0}, {-1, -1}, {-2, -2}, {-3, -3}},
	{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}},
	{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}},
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
}

var masOffsets = [][]Coord{
	{{1, 1}, {0, 0}, {-1, -1}},
	{{1, -1}, {0, 0}, {-1, 1}},
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

			for _, word := range getWords(chart, offsets, y, x) {
				if word == "XMAS" {
					xmasCount++
				}
			}
		}
	}

	fmt.Printf("XMAS: %d\n", xmasCount)

	masCount := 0
	for y, line := range chart {
	chars:
		for x, char := range line {
			if char != "A" {
				continue
			}
			words := getWords(chart, masOffsets, y, x)
			if len(words) != 2 {
				continue
			}

			for _, word := range words {
				if word != "MAS" && word != "SAM" {
					continue chars
				}
			}
			masCount++
		}
	}

	fmt.Printf("MAS: %d\n", masCount)
}

func getWords(chart [][]string, offsets [][]Coord, startY, startX int) []string {
	height := len(chart)
	width := len(chart[0])

	words := make([]string, 0, len(offsets))

offsets:
	for _, offset := range offsets {
		var word strings.Builder

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
