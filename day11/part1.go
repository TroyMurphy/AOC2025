package main

import (
	"fmt"
	"strings"
)

func Part1(lines []string) int {
	// inputs, _ := ParseCables(lines)
	_, outputs := ParseCables(lines)
	// fmt.Println(inputs["you"])
	// fmt.Println(outputs["you"])
	paths := PathsToOut(outputs, "you")
	return paths
}

func ParseCables(lines []string) (map[string][]string, map[string][]string) {
	inputs := make(map[string][]string)
	outputs := make(map[string][]string)
	for _, l := range lines {
		lineSplit := strings.Split(l, ":")
		device := lineSplit[0]
		deviceOutputs := strings.Trim(lineSplit[1], " ")
		deviceOut := strings.Split(deviceOutputs, " ")
		outputs[device] = deviceOut
		fmt.Println("Device ", device, " outputs to ", deviceOut)

		for _, o := range deviceOut {
			inputs[o] = append(inputs[o], device)
		}
	}
	return inputs, outputs
}

var foundPaths map[string]int = make(map[string]int)

func PathsToOut(outputs map[string][]string, node string) int {
	if node == "out" {
		return 1
	}
	pathsCount := 0
	for _, o := range outputs[node] {
		if val, exists := foundPaths[o]; exists {
			// fmt.Println("Found cached paths for ", o, ": ", val)
			pathsCount += val
			continue
		}
		pathsCount += PathsToOut(outputs, o)
	}
	foundPaths[node] = pathsCount
	// fmt.Println("Paths from ", node, " to out: ", path
	return pathsCount
}
