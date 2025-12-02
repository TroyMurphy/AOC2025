package main

import (
	"strconv"
	"strings"
)

func errh(e error) {
	if e != nil {
		panic(e)
	}
}

func Part2(lines []string) int64 {
	line := lines[0]
	ranges := strings.Split(line, ",")

	invalidSum := int64(0)

	for _, rangeString := range ranges {
		// rangeString is <int>-<int> parse them into min and max
		min, max, err := ParseRange(rangeString)
		errh(err)

		for _, i := range findInvalidIds(min, max) {
			invalidSum += i
		}
	}
	return invalidSum
}

func findInvalidIds(min, max int64) []int64 {
	invalidIds := []int64{}
	for search := min; search <= max; search++ {
		// check for regex match of a pure repeated pattern
		searchString := strconv.FormatInt(search, 10)

		if isRepeatedPattern(searchString) {
			invalidIds = append(invalidIds, search)
		}
	}
	return invalidIds
}

func isRepeatedPattern(s string) bool {
	n := len(s)
	for l := 1; l <= n/2; l++ {
		if n%l == 0 {
			pattern := s[:l]
			repeated := ""
			for i := 0; i < n/l; i++ {
				repeated += pattern
			}
			if repeated == s {
				return true
			}
		}
	}
	return false
}
