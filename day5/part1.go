package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type freshRange struct {
	min, max int64
}

func Parse(lines []string) ([]freshRange, []int64) {
	var ingredientIds []int64

	var ranges []freshRange
	parseIngredients := false

	for _, line := range lines {
		if line == "" {
			parseIngredients = true
			continue
		}
		if parseIngredients {
			ingredientID, err := strconv.ParseInt(line, 10, 64)
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
		min, _ := strconv.ParseInt(parts[0], 10, 64)
		max, _ := strconv.ParseInt(parts[1], 10, 64)
		ranges = append(ranges, freshRange{min: min, max: max})
	}
	return ranges, ingredientIds
}

func Part1(lines []string) int64 {
	return CountFresh(Parse(lines))
}

func Part2(lines []string) int64 {
	ranges, _ := Parse(lines)
	return CountRanges(CollapseRanges(ranges))
}

func CountFresh(ranges []freshRange, ingredients []int64) int64 {
	outCount := int64(0)
	for _, i := range ingredients {
		if IsFresh(ranges, i) {
			outCount++
		}
	}
	return outCount
}

func IsFresh(ranges []freshRange, ingredient int64) bool {
	for _, r := range ranges {
		if ingredient >= r.min && ingredient <= r.max {
			return true
		}
	}
	return false
}

func CollapseRanges(inRanges []freshRange) []freshRange {
	sort.Slice(inRanges, func(i, j int) bool {
		return inRanges[i].min < inRanges[j].min
	})

	if len(inRanges) == 0 {
		return []freshRange{}
	}
	target := inRanges[0]
	return CollapseRangesRec(inRanges[1:], []freshRange{}, target.min, target.max)
}

func CollapseRangesRec(inRanges, outRanges []freshRange, trackedMin, trackedMax int64) []freshRange {
	// fmt.Println(outRanges)
	if len(inRanges) == 0 {
		outRanges = append(outRanges, freshRange{min: trackedMin, max: trackedMax})
		// fmt.Println(outRanges)
		// fmt.Println("Done")
		return outRanges
	}
	target := inRanges[0]
	// fmt.Printf("Tracked: %d->%d | Target: %d->%d\r\n", trackedMin, trackedMax, target.min, target.max)
	if trackedMax < target.min {
		newOut := append(outRanges, freshRange{min: trackedMin, max: trackedMax})
		return CollapseRangesRec(inRanges[1:], newOut, target.min, target.max)
	}
	return CollapseRangesRec(inRanges[1:], outRanges, min(trackedMin, target.min), max(trackedMax, target.max))
}

func CountRanges(ranges []freshRange) int64 {
	outCount := int64(0)
	for _, r := range ranges {
		outCount += r.max - r.min + int64(1)
	}
	return outCount
}
