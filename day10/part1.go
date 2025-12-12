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

var (
	logDebug bool = true
	logHelp  bool = true
)

func Part1(lines []string) int64 {
	machines := ParseMachines(lines)

	presses := int64(0)
	for _, m := range machines {
		minpresses := MinPresses([]clickState{
			{lights: m.lights, buttons: m.buttons, clicked: 0},
		})
		if logHelp {
			fmt.Printf("Machine lights: %0b requires min presses: %d\r\n", m.lights, minpresses)
		}
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
	lightsMatch = strings.ReplaceAll(lightsMatch, ".", "0")
	lightsMatch = strings.ReplaceAll(lightsMatch, "#", "1")
	lights, _ := strconv.ParseInt(lightsMatch, 2, 32)

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
	if logDebug {
		fmt.Printf("Parsed button: %s -> %0b\r\n", match, output)
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

func MinPresses(states []clickState) int64 {
	if len(states) == 0 {
		return -1
	}
	for _, s := range states {
		if !IsPossible(s) {
			continue
		}
		if logDebug {
			fmt.Printf("Current lights: %0b\r\n", s.lights)
		}
		toClick := s.buttons
		for i, b := range toClick {
			afterClick := s.lights ^ b
			if logDebug {
				fmt.Printf("  Trying button: %0b -> %0b\r\n", b, afterClick)
			}
			if afterClick == 0 {
				return s.clicked + 1
			}
			paredButtons := slices.Clone(toClick)
			paredButtons = slices.Delete(paredButtons, i, i+1)
			if len(paredButtons) == 1 {
				if paredButtons[0] == s.lights {
					return s.clicked + 2
				}
				continue
			}
			if len(paredButtons) > 0 {
				newState := clickState{lights: afterClick, buttons: paredButtons, clicked: s.clicked + 1}
				newStates = append(newStates, newState)
			}
		}
	}
	return MinPresses(newStates)
}

func SerializeState(state clickState) string {
	// format the light, then the ordered buttons
	return fmt.Sprintf("%0b|%v", state.lights, state.buttons)
}

func IsPossible(state clickState) bool {
	lights := state.lights
	buttons := state.buttons
	maxBit := 0
	for temp := lights; temp > 0; temp >>= 1 {
		maxBit++
	}
	for bit := 0; bit < maxBit; bit++ {
		bitmask := int64(1) << bit
		if (lights & bitmask) != 0 {
			bitFound := false
			for _, b := range buttons {
				if (b & bitmask) != 0 {
					bitFound = true
					break
				}
			}
			if !bitFound {
				return false
			}
		}
	}
	return true
}
