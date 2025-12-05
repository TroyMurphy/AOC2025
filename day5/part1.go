package main

import (
	"fmt"
	"strconv"
	"strings"
)

type freshRange struct {
	min, max int
}

func Parse(lines []string) ([]freshRange, []int) {
	var ingredientIds []int

	var ranges []freshRange
	parseIngredients := false

	for _, line := range lines {
		if line == "" {
			parseIngredients = true
			continue
		}
		if parseIngredients {
			ingredientID, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ingredientIds = append(ingredientIds, ingredientID)
			continue
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			fmt.Printf("Erorr parsing %s", line)
		}
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, freshRange{min: min, max: max})
	}
	return ranges, ingredientIds
}

func Part1(lines []string) int {
	return CountFresh(Parse(lines))
}

func Part2(lines []string) int {
	ranges, _ := Parse(lines)
	return CountRanges(CollapseRangesRec(ranges, []freshRange{}))
}

func CountFresh(ranges []freshRange, ingredients []int) int {
	outCount := 0
	for _, i := range ingredients {
		if IsFresh(ranges, i) {
			outCount++
		}
	}
	return outCount
}

func IsFresh(ranges []freshRange, ingredient int) bool {
	for _, r := range ranges {
		if ingredient >= r.min && ingredient <= r.max {
			return true
		}
	}
	return false
}

// func CollapseRangesRec(inRanges []freshRange, outRanges []freshRange) []freshRange {
// 	if len(inRanges) == 0 {
// 		return outRanges
// 	}
//
// 	insert := inRanges[0]
// 	var extendedRanges []freshRange
//
// 	extendedLow := false
// 	extendedHigh := false
// 	for _, or := range outRanges {
// 		if insert.min >= or.min && insert.min <= or.max {
// 			extendedRanges = append(extendedRanges, freshRange{min: or.min, max: max(or.max, insert.max)})
// 			extendedHigh = true
// 			break
// 		}
// 		if insert.max >= or.min && insert.max <= or.max {
// 			extendedRanges = append(extendedRanges, freshRange{min: min(or.min, insert.min), max: or.max})
// 			extendedLow = true
// 		}
// 	}
// 	if !extendedLow && !extendedHigh {
// 		extendedRanges = append(extendedRanges, insert)
// 	}
// 	// this doesn't work
//
// }

func CountRanges(ranges []freshRange) int {
	outCount := 0
	for _, r := range ranges {
		outCount += r.max - r.min + 1
	}
	return outCount
}
