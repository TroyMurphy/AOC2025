package main

import "fmt"

func Part2(lines []string) int {
	boxes := Parse(lines)
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

	sortedPairs := SortMap(distanceMap)

	box1, box2 := BuildCircuitMapEndless(sortedPairs, len(lines))

	output := box1.x * box2.x
	return output
}

func BuildCircuitMapEndless(sortedPairs []boxPair, maxBoxes int) (box, box) {
	boxToCircuit := make(map[box]int)
	circuitToBoxes := make(map[int][]box)

	var box1 box
	var box2 box

	i := 0
	for _, boxes := range sortedPairs {
		// check the first key of circuitToBoxes(since we are cleaning up)
		// if it contains the same size as all boxes, we quit.
		for _, boxes := range circuitToBoxes {
			if len(boxes) == maxBoxes {
				return box1, box2
			}
			break
		}

		box1 = boxes.box1
		box2 = boxes.box2
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
			circuitKey := min(box1Circuit, box2Circuit)
			otherKey := max(box1Circuit, box2Circuit)
			if circuitKey == otherKey {
				i++
				continue
			}
			otherCircuitBoxes := circuitToBoxes[otherKey]
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
			i++
			continue
		}
		if box2Exists {
			boxToCircuit[boxes.box1] = box2Circuit
			circuitToBoxes[box2Circuit] = append(circuitToBoxes[box2Circuit], boxes.box1)
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
	panic("Circuit didn't close")
}
