package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed first_input
var content []byte

func main() {
	left := make([]int, 0)
	right := make([]int, 0)

	lines := bytes.Split(content, []byte("\n"))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(string(line), "   ")
		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		left = append(left, leftNum)

		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Printf("Distance: %d\n", distance)
}
