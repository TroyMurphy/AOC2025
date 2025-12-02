package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func Part2(lines []string) int {
	zeroCount := 0
	current := 50
	for _, line := range lines {
		parsed := strings.Replace(line, "R", "", 1)
		parsed = strings.Replace(parsed, "L", "-", 1)
		turn, err := strconv.Atoi(parsed)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Current: %d -> %d\r\n", current, turn)

		fullRotations := int(math.Floor(math.Abs(float64(turn)) / 100))
		remainder := turn % 100
		end := current
		zeroCount += fullRotations

		if remainder == 0 {
			fmt.Println("Remainder is 0")
			continue
		}

		if remainder < 0 {
			if current+remainder <= 0 && current != 0 {
				zeroCount += 1
				fmt.Printf("Left past 0 <%d>\r\n", zeroCount)
			}
			end = mod(current+remainder, 100)
			fmt.Printf("End on %d\r\n", end)
			current = end
			continue
		}
		// if remainder > 0 {
		if current+remainder >= 100 {
			zeroCount += 1
			fmt.Printf("Right past 0 <%d> \r\n ", zeroCount)
		}
		end = mod(current+remainder, 100)
		fmt.Printf("End on %d\r\n", end)
		current = end
		continue

	}
	return zeroCount
}
