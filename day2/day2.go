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
	res, _ := checkLevels(r.levels)
	return res
}

func (r *Report) MostlySafe() bool {
	res, failedIndex := checkLevels(r.levels)
	if res {
		return true
	}

	if failedIndex == 2 {
		if res, _ := checkLevels(slices.Delete(slices.Clone(r.levels), 0, 1)); res {
			return true
		}
	}
	if res, _ := checkLevels(slices.Delete(slices.Clone(r.levels), failedIndex-1, failedIndex)); res {
		return true
	}
	if res, _ := checkLevels(slices.Delete(slices.Clone(r.levels), failedIndex, failedIndex+1)); res {
		return true
	}

	return false
}

func checkLevels(levels []int) (bool, int) {
	last := levels[0]
	desc := levels[0] > levels[1]

	for i, level := range levels[1:] {
		diff := last - level

		if desc && (diff < 1 || diff > 3) {
			return false, i + 1
		} else if !desc && (diff > -1 || diff < -3) {
			return false, i + 1
		}

		last = level
	}

	return true, 0
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
