package main

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}
type pair struct {
	p1, p2 point
}

func Part1(lines []string) int64 {
	points := ParsePoints(lines)
	pairs := CreatePairs(points)
	fmt.Printf("%d pairs\r\n", len(pairs))
	maxArea := int64(0)
	for _, p := range pairs {
		area := Area(p)
		maxArea = max(area, maxArea)
		// if area != maxArea {
		// 	fmt.Printf("(%d,%d), (%d, %d)\r\n", p.p1.x, p.p1.y, p.p2.x, p.p2.y)
		// }
	}
	return maxArea
}

func Area(p pair) int64 {
	return int64(abs(p.p1.x-p.p2.x)+1) * int64(abs(p.p1.y-p.p2.y)+1)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func ParsePoints(lines []string) []point {
	var points []point
	for _, l := range lines {
		values := strings.Split(l, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		points = append(points, point{x: x, y: y})
	}
	return points
}

func CreatePairs(points []point) []pair {
	var pairs []pair
	for i, p := range points {
		for j, p2 := range points {
			if i < j {
				pairs = append(pairs, pair{p1: p, p2: p2})
			}
		}
	}
	return pairs
}
