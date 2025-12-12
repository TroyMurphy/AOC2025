package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type jmachine struct {
	buttons  [][]int64
	joltages []int64
}

func Part2(lines []string) int64 {
	jmachines := ParseJMachines(lines)
	// fmt.Println(jmachines[0])
	presses := int64(0)
	for _, j := range jmachines {
		presses += MinJPresses(j)
	}
	return presses
}

func MinJPresses(jmachine jmachine) int64 {
	return 0
}

func ParseJMachines(lines []string) []jmachine {
	var machines []jmachine
	for _, l := range lines {
		machines = append(machines, ParseJMachine(l))
	}
	return machines
}

func ParseJMachine(line string) jmachine {
	joltagesre := regexp.MustCompile(`\{.*\}`)
	joltagesMatch := joltagesre.FindAllString(line, -1)[0]
	joltagesMatch = strings.ReplaceAll(joltagesMatch, "{", "")
	joltagesMatch = strings.ReplaceAll(joltagesMatch, "}", "")
	joltageStrings := strings.Split(joltagesMatch, ",")
	var joltages []int64
	for _, j := range joltageStrings {
		value, _ := strconv.ParseInt(j, 10, 64)
		joltages = append(joltages, value)
	}
	matrixLength := len(joltages)

	buttonre := regexp.MustCompile(`\(.*\)`)
	buttonMatch := buttonre.FindAllString(line, -1)[0]
	buttonStrings := strings.Split(buttonMatch, " ")
	var buttons [][]int64
	for _, b := range buttonStrings {
		buttons = append(buttons, ParseMatrixButton(b, matrixLength))
	}

	return jmachine{buttons: buttons, joltages: joltages}
}

func ParseMatrixButton(buttonString string, length int) []int64 {
	buttonFullString := strings.ReplaceAll(buttonString, "(", "")
	buttonFullString = strings.ReplaceAll(buttonFullString, ")", "")
	buttonStringValues := strings.Split(buttonFullString, ",")
	// fmt.Println("Button String Values: ", buttonStringValues)
	var buttonIndexes []int
	for _, bs := range buttonStringValues {
		bi, _ := strconv.Atoi(bs)
		buttonIndexes = append(buttonIndexes, bi)
	}

	var output []int64

	for i := range length {
		if slices.Contains(buttonIndexes, i) {
			output = append(output, int64(1))
			continue
		}
		output = append(output, int64(0))
	}
	return output
}
