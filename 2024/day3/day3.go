package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var mulRegex = regexp.MustCompile(`mul\((\d+,\d+)\)`)

//go:embed input
var input string

func main() {
	total := 0

	muls := mulRegex.FindAllStringSubmatch(input, -1)
	for _, mul := range muls {
		nums := strings.Split(mul[1], ",")
		fmt.Println(nums[1])
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

	fmt.Printf("Total %d\n", total)
}
