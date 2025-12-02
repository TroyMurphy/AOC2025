package main

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	zeroCount := 0
	current := 50
	for _, line := range lines {
		parsed := strings.Replace(line, "R", "", 1)
		parsed = strings.Replace(parsed, "L", "-", 1)
		turn, err := strconv.Atoi(parsed)
		if err != nil {
			panic(err)
		}

		current = (current + turn) % 100
		if current == 0 {
			zeroCount += 1
		}
	}
	return zeroCount
}
