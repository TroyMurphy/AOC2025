package main

import (
	"strconv"
)

func Part2(lines []string) int64 {
	voltage := int64(0)

	for _, line := range lines {
		lineMax := MaxVoltage(line)
		// fmt.Printf("%d\r\n", lineMax)
		voltage += lineMax
	}
	return voltage
}

func MaxVoltage(line string) int64 {
	var numbers []int

	for _, c := range line {
		if val, err := strconv.Atoi(string(c)); err == nil {
			numbers = append(numbers, val)
		}
	}
	return MaxVoltageRec(numbers, 12, 0)
}

func MaxVoltageRec(numbers []int, take int, acc int64) int64 {
	if take == 0 {
		return acc
	}

	maxIndex := len(numbers) - take + 1
	available := numbers[0:maxIndex]
	// fmt.Printf("%v\r\n", available)

	searchIndex := 0
	searchMax := available[0]
	for i, x := range available[1:] {
		if x > searchMax {
			searchMax = x
			searchIndex = i + 1
			// fmt.Printf("New max: %d at %d", searchMax, searchIndex)
		}
	}

	return MaxVoltageRec(numbers[searchIndex+1:], take-1, (acc*int64(10))+int64(searchMax))
}
