package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	levels := make([]int, 0)

	for _, rawLevel := range rawLevels {
		level, err := strconv.Atoi(rawLevel)
		if err != nil {
			panic(err)
		}
		levels = append(levels, level)
	}

	return Report{
		levels: levels,
	}
}

func (r *Report) Safe() bool {
	return checkLevels(r.levels, false)
}

func (r *Report) MostlySafe() bool {
	return checkLevels(r.levels, true)
}

func checkLevels(levels []int, singleException bool) bool {
	last := levels[0]
	desc := levels[0] > levels[1]

	for i, level := range levels[1:] {
		diff := last - level
		safe := true

		if desc && (diff < 1 || diff > 3) {
			safe = false
		} else if !desc && (diff > -1 || diff < -3) {
			safe = false
		}

		if !safe && singleException {
			n := i + 1

			pre := slices.Delete(slices.Clone(levels), n, n+1)
			post := slices.Delete(slices.Clone(levels), n-1, n)
			firstRes := false

			// edge case: handle removing the first element
			if n == 2 {
				firstRes = checkLevels(slices.Delete(slices.Clone(levels), 0, 1), false)
			}

			return firstRes || checkLevels(pre, false) || checkLevels(post, false)
		} else if !safe {
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

	safe = 0
	for _, line := range reports {
		report := NewReport(line)

		if report.MostlySafe() {
			safe++
		}
	}

	fmt.Printf("Safe reports with 1 exception: %d\n", safe)
}
