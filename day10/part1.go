package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type lmachine struct {
	lights  int64
	buttons []int64
}

func Part1(lines []string) int64 {
	machines := ParseMachines(lines)

	presses := int64(0)
	for _, m := range machines {
		fmt.Printf("Target: %b\r\n", m.lights)
		minpresses := MinPresses(m.lights, []reachableState{{value: 0, remainingButtons: m.buttons}}, int64(0))
		presses += minpresses
	}
	return presses
}

func ParseMachines(lines []string) []lmachine {
	var machines []lmachine
	for _, l := range lines {
		machines = append(machines, ParseMachine(l))
	}
	return machines
}

func ParseMachine(line string) lmachine {
	lightsre := regexp.MustCompile(`\[.*\]`)
	lightsMatch := lightsre.FindAllString(line, -1)[0]
	lightsMatch = strings.ReplaceAll(lightsMatch, "[", "")
	lightsMatch = strings.ReplaceAll(lightsMatch, "]", "")
	lights := LightsToInt(lightsMatch)
	fmt.Printf("Lights: %b \r\n", lights)

	var buttons []int64
	buttonre := regexp.MustCompile(`\(.*\)`)
	buttonMatch := buttonre.FindAllString(line, -1)[0]
	buttonStrings := strings.Split(buttonMatch, " ")
	for _, match := range buttonStrings {
		buttons = append(buttons, ParseButton(match))
	}

	return lmachine{lights: lights, buttons: buttons}
}

func ParseButton(match string) int64 {
	match = strings.ReplaceAll(match, "(", "")
	match = strings.ReplaceAll(match, ")", "")
	values := strings.Split(match, ",")
	output := int64(0)
	for _, v := range values {
		vint, _ := strconv.ParseInt(v, 10, 32)
		output += twoPow(vint)
	}
	return output
}

func twoPow(exp int64) int64 {
	return int64(math.Pow(float64(2), float64(exp)))
}

type clickState struct {
	lights  int64
	buttons []int64
	clicked int64
}

var newStates []clickState = []clickState{}

type reachableState struct {
	value            int64
	remainingButtons []int64
}

var checked map[string]struct{} = make(map[string]struct{})

func MinPresses(target int64, reachable []reachableState, acc int64) int64 {
	if len(reachable) == 0 {
		return -1
	}
	fmt.Printf("Searching %d states\r\n", len(reachable))
	var validStates []reachableState
	for _, s := range reachable {
		if s.value == target {
			return acc
		}
		if len(s.remainingButtons) > 0 {
			serialized := Serialize(&s)
			if _, exists := checked[serialized]; exists {
				fmt.Printf("Skipping already checked state: %s\r\n", serialized)
				continue
			}
			validStates = append(validStates, s)
			checked[serialized] = struct{}{}
		}
	}
	fmt.Printf("%d valid states\r\n", len(validStates))
	var newStates []reachableState
	for _, s := range validStates {
		fmt.Printf("Start: %b | Buttons: ", s.value)
		for bi, b := range s.remainingButtons {
			fmt.Printf("%b, ", b)
			reached := s.value ^ b
			remaining := slices.Delete(slices.Clone(s.remainingButtons), bi, bi+1)
			newState := reachableState{value: reached, remainingButtons: remaining}
			newStates = append(newStates, newState)
		}
		fmt.Printf("\r\n")
	}
	// buttons should be reduced to
	fmt.Printf("Generated %d new states\r\n", len(newStates))
	return MinPresses(target, newStates, acc+1)
}

func Serialize(state *reachableState) string {
	return fmt.Sprintf("%d|%v", state.value, state.remainingButtons)
}

func LightsToInt(lights string) int64 {
	output := int64(0)
	for i := 0; i < len(lights); i++ {
		if lights[i] == '#' {
			output += twoPow(int64(i))
		}
	}
	return output
}
