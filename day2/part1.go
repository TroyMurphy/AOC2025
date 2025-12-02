package main

import (
	"day2/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ParseRange(s string) (min, max int64, err error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid range format")
	}
	min, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	max, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return min, max, nil
}

func Part1(lines []string) int64 {
	line := lines[0]
	ranges := strings.Split(line, ",")

	invalidSum := int64(0)

	for _, rangeString := range ranges {
		// rangeString is <int>-<int> parse them into min and max
		min, max, err := ParseRange(rangeString)
		utils.Check(err)

		for _, i := range InvalidIds(min, max) {
			invalidSum += i
		}
	}

	return invalidSum
}

func InvalidIds(min, max int64) []int64 {
	stringMin := strconv.FormatInt(min, 10)
	var badIds []int64

	// fmt.Printf("Min: %s, Max: %s\r\n", stringMin, stringMax)
	if IsInvalid(stringMin) {
		badIds = append(badIds, min)
	}
	search := NextSearch(min)
	for search <= max {
		badIds = append(badIds, search)
		search = NextSearch(search)
	}

	return badIds
}

func IsInvalid(value string) bool {
	if value[0] == '0' {
		return false
	}
	if len(value)%2 != 0 {
		return false
	}
	half := len(value) / 2
	return value[0:half] == value[half:]
}

func NextSearch(intVal int64) int64 {
	value := strconv.FormatInt(intVal, 10)

	if len(value)%2 != 0 {
		// not symmetric. Therefore, next invalid number
		zeroPadding := math.Floor(float64(len(value)) / 2)
		nextStartString := "1"
		nextStartString += strings.Repeat("0", int(zeroPadding))
		output, err := strconv.ParseInt(strings.Repeat(nextStartString, 2), 10, 64)
		utils.Check(err)
		return output
	}
	// it is symmetric
	startString := value[0:int(len(value)/2)]
	if startString+startString > value {
		output, err := strconv.ParseInt(strings.Repeat(startString, 2), 10, 64)
		utils.Check(err)
		return output
	}

	startNum, err := strconv.ParseInt(startString, 10, 64)
	utils.Check(err)
	nextStartString := strconv.FormatInt(startNum+1, 10)
	utils.Check(err)
	output, err := strconv.ParseInt(strings.Repeat(nextStartString, 2), 10, 64)
	utils.Check(err)
	return output
}
