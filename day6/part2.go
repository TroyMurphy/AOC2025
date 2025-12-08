package main

import (
	"strconv"
	"strings"
)

func Part2(lines []string) int64 {
	// values := ParseNums(lines[0 : len(lines)-1])
	// return Solve1(values, operators, 0)
	nums := ParseNums(lines[0 : len(lines)-1])
	ops := ParseOps(lines[len(lines)-1])
	// fmt.Println(nums)

	return Zip(nums, ops, int64(0))
}

func ParseNums(lines []string) [][]int64 {
	index := 0
	var numberArrays [][]int64
	var numberArray []int64
	for {
		if index >= len(lines[0]) {
			numberArrays = append(numberArrays, numberArray)
			break
		}
		var b strings.Builder
		for _, l := range lines {
			b.WriteByte(l[index])
		}
		colVal := strings.Trim(b.String(), " ")
		// fmt.Println(colVal)
		if colVal == "" {
			numberArrays = append(numberArrays, numberArray)
			numberArray = nil
			index++
			continue
		}
		colNum, _ := strconv.ParseInt(colVal, 10, 64)
		numberArray = append(numberArray, colNum)
		// fmt.Println(numberArray)
		index++
	}

	return numberArrays
}

func Zip(ngs [][]int64, ops []string, acc int64) int64 {
	if len(ngs) == 0 {
		return acc
	}

	ng := ngs[0]
	op := ops[0]

	if op == "*" {
		product := int64(1)
		for _, v := range ng {
			product *= v
		}
		return Zip(ngs[1:], ops[1:], acc+product)
	}
	sum := int64(0)
	for _, v := range ng {
		sum += v
	}
	return Zip(ngs[1:], ops[1:], acc+sum)
}
