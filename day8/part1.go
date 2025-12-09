package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type box struct {
	x, y, z int
}
type boxPair struct {
	box1, box2 box
}

var log bool = false

func Part1(lines []string) int {
	boxes := Parse(lines)
	if log {
		fmt.Printf("Read %d Junction Boxes\r\n", len(boxes))
	}
	distanceMap := make(map[boxPair]float64)

	for i, box1 := range boxes {
		for j, box2 := range boxes {
			if j <= i {
				continue
			}
			distance := Distance(box1, box2)
			pair := boxPair{box1: box1, box2: box2}
			distanceMap[pair] = distance
		}
	}

	if log {
		fmt.Printf("Created %d pairs\r\n", len(distanceMap))
	}

	sortedPairs := SortMap(distanceMap)

	circuitMap := BuildCircuitMap(sortedPairs, 1000)
	// for k := range circuitMap {
	// 	// output max circuits
	// 	fmt.Printf("Circuit %d -> [%d]\r\n", k, len(circuitMap[k])
	// }
	// if log {
	// 	for cmk := range circuitMap {
	// 		fmt.Printf("Circuit %d: %d\r\n", cmk, len(circuitMap[cmk]))
	// 	}
	// }
	topCircuitSizes := TopCircuitSizes(circuitMap, 3)
	output := 1
	for _, cs := range topCircuitSizes {
		output *= cs
	}
	return output
}

func BuildCircuitMap(sortedPairs []boxPair, iterations int) map[int][]box {
	boxToCircuit := make(map[box]int)
	circuitToBoxes := make(map[int][]box)

	i := 0
	for _, boxes := range sortedPairs {
		if i >= iterations {
			break
		}
		box1 := boxes.box1
		box2 := boxes.box2
		if log {
			fmt.Printf("Box1: {%d, %d, %d} | Box2: {%d, %d, %d}\r\n", box1.x, box1.y, box1.z, box2.x, box2.y, box2.z)
		}
		// if box1 exists in boxToCircuit
		//   if box2 exists in boxToCircuit
		//     get all boxes in circuitToBox from boxToCircuit[box2]
		//   else
		//     set box2 to boxToCircuit[box1]
		// else if box2 exists in boxToCircuit
		//   same thing as box 1
		// If neither of them exist then we use the current i index as the circuit for both boxes
		box1Circuit, box1Exists := boxToCircuit[boxes.box1]
		box2Circuit, box2Exists := boxToCircuit[boxes.box2]

		if box1Exists && box2Exists {
			if log {
				fmt.Printf("Conflict between circuit %d and %d\r\n", box1Circuit, box2Circuit)
			}
			circuitKey := min(box1Circuit, box2Circuit)
			otherKey := max(box1Circuit, box2Circuit)
			if circuitKey == otherKey {
				if log {
					fmt.Println("Same circuit already. Skipping")
				}
				i++
				continue
			}
			otherCircuitBoxes := circuitToBoxes[otherKey]
			if log {
				fmt.Printf("Moving %d boxes from circuit %d into circuit %d\r\n", len(otherCircuitBoxes), otherKey, circuitKey)
			}
			for _, ocb := range otherCircuitBoxes {
				boxToCircuit[ocb] = circuitKey
				circuitToBoxes[circuitKey] = append(circuitToBoxes[circuitKey], ocb)
			}
			delete(circuitToBoxes, otherKey)
			i++
			continue
		}
		if box1Exists {
			boxToCircuit[boxes.box2] = box1Circuit
			circuitToBoxes[box1Circuit] = append(circuitToBoxes[box1Circuit], boxes.box2)
			if log {
				fmt.Printf("Added box {%d, %d, %d} to circuit %d. Now size %d\r\n", box2.x, box2.y, box2.z, box1Circuit, len(circuitToBoxes[box1Circuit]))
			}
			i++
			continue
		}
		if box2Exists {
			boxToCircuit[boxes.box1] = box2Circuit
			circuitToBoxes[box2Circuit] = append(circuitToBoxes[box2Circuit], boxes.box1)
			if log {
				fmt.Printf("Added box {%d, %d, %d} to circuit %d. Now size %d\r\n", box1.x, box1.y, box1.z, box2Circuit, len(circuitToBoxes[box2Circuit]))
			}
			i++
			continue
		}
		// default
		circuitToBoxes[i] = []box{box1, box2}
		boxToCircuit[box1] = i
		boxToCircuit[box2] = i
		if log {
			fmt.Printf("Added box {%d, %d, %d} and box {%d, %d, %d} to circuit %d\r\n", box1.x, box1.y, box1.z, box2.x, box2.y, box2.z, i)
		}
		i++
		continue
	}
	return circuitToBoxes
}

func Parse(lines []string) []box {
	var output []box
	for _, line := range lines {

		values := strings.Split(line, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])
		output = append(output, box{x: x, y: y, z: z})
	}
	return output
}

func Distance(p1 box, p2 box) float64 {
	base := math.Pow(float64(p1.x-p2.x), float64(2)) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2)
	return math.Sqrt(base)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func SortMap(pairDistances map[boxPair]float64) []boxPair {
	keys := make([]boxPair, 0, len(pairDistances))
	for key := range pairDistances {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool { return pairDistances[keys[i]] < pairDistances[keys[j]] })

	return keys
}

func TopCircuitSizes(circuits map[int][]box, take int) []int {
	var sizes []int

	for k := range circuits {
		sizes = append(sizes, len(circuits[k]))
	}

	sort.Slice(sizes, func(i int, j int) bool { return sizes[i] > sizes[j] })
	return sizes[0:take]
}
