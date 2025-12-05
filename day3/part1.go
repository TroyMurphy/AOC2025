package main

import (
	"strconv"
)

func Part1(lines []string) int {
	sum := 0
	for _, val := range lines {
		sum += MaxVal(val)
	}
	return sum
}

func MaxVal(value string) int {
	tensMax := int64(0)
	tensIndex := 0
	searchIndex := 0

	for searchIndex < len(value)-1 {
		current, _ := strconv.ParseInt(string(value[searchIndex]), 10, 64)

		if current > int64(tensMax) {
			tensMax = current
			tensIndex = searchIndex
		}
		searchIndex++
	}

	onesMax := int64(0)
	searchIndex = tensIndex + 1
	for searchIndex < len(value) {
		current, _ := strconv.Atoi(string(value[searchIndex]))

		if current > int(onesMax) {
			onesMax = int64(current)
		}
		searchIndex++
	}

	maxVal, err := strconv.Atoi(strconv.FormatInt(tensMax, 10) + strconv.FormatInt(onesMax, 10))
	if err != nil {
		panic(err)
	}
	return maxVal
}
