package main

import (
	"regexp"
	"strconv"
)

func Part1(lines []string) int64 {
	numLists := ParseLists(lines, [][]int64{})
	operators := ParseOps(lines[len(lines)-1])
	// fmt.Println(numLists)
	// fmt.Println(operators)

	return Solve1(numLists, operators, 0)
}

func ParseLists(lines []string, acc [][]int64) []([]int64) {
	if len(lines) == 1 {
		return acc
	}
	numRegex := regexp.MustCompile(`\d+`)
	var output []int64
	numbers := numRegex.FindAllString(lines[0], -1)
	for _, n := range numbers {
		val, _ := strconv.ParseInt(n, 10, 64)
		output = append(output, val)
	}
	acc = append(acc, output)
	return ParseLists(lines[1:], acc)
}

func ParseOps(line string) []string {
	opReex := regexp.MustCompile(`\+|\*`)
	return opReex.FindAllString(line, -1)
}

func Solve1(numLists [][]int64, operators []string, index int) int64 {
	if index >= len(operators) {
		return 0
	}

	isMultiply := operators[index] == "*"

	column := int64(0)
	if isMultiply {
		column = int64(1)
	}
	for _, x := range numLists {
		if isMultiply {
			column *= x[index]
			continue
		}
		column += x[index]
	}

	return column + Solve1(numLists, operators, index+1)
}
