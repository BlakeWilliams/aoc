package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var mulRegex = regexp.MustCompile(`mul\((\d+,\d+)\)|don't\(\)|do\(\)`)

//go:embed input
var input string

func main() {
	part1()
	part2()
}

func part1() {
	total := 0

	muls := mulRegex.FindAllStringSubmatch(input, -1)
	for _, mul := range muls {
		if mul[0] == "do()" || mul[0] == "don't()" {
			continue
		}

		nums := strings.Split(mul[1], ",")
		first, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		total += first * second
	}

	fmt.Printf("Total: %d\n", total)
}

func part2() {
	total := 0
	muls := mulRegex.FindAllStringSubmatch(input, -1)
	enabled := true

	for _, mul := range muls {
		if mul[0] == "don't()" {
			enabled = false
			continue
		} else if mul[0] == "do()" {
			enabled = true
			continue
		}

		if !enabled {
			continue
		}

		nums := strings.Split(mul[1], ",")
		first, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		total += first * second
	}

	fmt.Printf("Total with toggle: %d\n", total)
}
