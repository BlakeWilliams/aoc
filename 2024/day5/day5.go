package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := strings.Split(parts[0], "\n")
	updates := strings.Split(parts[1], "\n")

	ruleLookup := make(map[int][]int, len(rules))

	for _, rawRule := range rules {
		parts := strings.Split(rawRule, "|")

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		if _, ok := ruleLookup[left]; !ok {
			ruleLookup[left] = make([]int, 0, 1)
		}

		ruleLookup[left] = append(ruleLookup[left], right)
	}

	sum := 0
updates:
	for _, update := range updates {
		values := updateToInts(update)

		for i := len(values) - 1; i >= 0; i-- {
			rules, ok := ruleLookup[values[i]]
			if !ok {
				continue
			}

			for n := i; n >= 0; n-- {
				for _, rule := range rules {
					if values[n] == rule {
						continue updates
					}
				}
			}
		}

		mid := values[len(values)/2]
		sum += mid
	}

	fmt.Printf("Sum: %d\n", sum)
}

func updateToInts(s string) []int {
	numbers := strings.Split(s, ",")
	res := make([]int, len(numbers))
	for i, rawNum := range numbers {
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			panic(err)
		}

		res[i] = num
	}

	return res
}
