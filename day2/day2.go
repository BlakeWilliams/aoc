package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

type Report struct {
	levels []int
}

func NewReport(line string) Report {
	rawLevels := strings.Split(line, " ")
	levels := make([]int, len(rawLevels))

	for i, rawLevel := range rawLevels {
		level, err := strconv.Atoi(rawLevel)
		if err != nil {
			panic(err)
		}
		levels[i] = level
	}

	return Report{
		levels: levels,
	}
}

func (r *Report) Safe() bool {
	last := r.levels[0]
	desc := r.levels[0] > r.levels[1]

	for _, level := range r.levels[1:] {
		diff := last - level
		if desc && (diff < 1 || diff > 3) {
			return false
		} else if !desc && (diff > -1 || diff < -3) {
			return false
		}
		last = level
	}

	return true
}

func main() {
	reports := strings.Split(strings.TrimSpace(input), "\n")

	safe := 0
	for _, line := range reports {
		report := NewReport(line)

		if report.Safe() {
			safe++
		}
	}

	fmt.Printf("Safe reports: %d\n", safe)
}
