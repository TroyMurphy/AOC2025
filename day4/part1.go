package main

import (
	"slices"
)

type point struct {
	x, y int
}

func Part1(lines []string) int {
	var points []point
	for ii, ix := range lines {
		for ji, jx := range ix {
			if jx == '@' {
				newPoint := point{x: ii, y: ji}
				points = append(points, newPoint)
			}
		}
	}
	// fmt.Println(points)

	return CountAccessible(points)
}

func CountAccessible(points []point) int {
	count := 0
	for _, p := range points {
		if CanAccess(p, points, 4) {
			count++
		}
	}
	return count
}

func CanAccess(p point, points []point, threshold int) bool {
	diffs := [3]int{-1, 0, 1}
	// fmt.Printf("Searching around {%d, %d}\r\n", p.x, p.y)

	var searchPoints []point

	for _, dx := range diffs {
		for _, dy := range diffs {
			if dx == 0 && dy == 0 {
				continue
			}
			searchPoints = append(searchPoints, point{x: p.x + dx, y: p.y + dy})
		}
	}

	// fmt.Println(searchPoints)

	occupied := 0
	for _, search := range searchPoints {
		if slices.Contains(points, search) {
			occupied++
		}
	}
	// fmt.Printf("Found %d\r\n", occupied)
	return occupied < threshold
}

func Part2(lines []string) int {
	var points []point
	for ii, ix := range lines {
		for ji, jx := range ix {
			if jx == '@' {
				newPoint := point{x: ii, y: ji}
				points = append(points, newPoint)
			}
		}
	}
	return CountAccessibleRec(points, 0)
}

func CountAccessibleRec(points []point, removed int) int {
	var accessible []point
	for _, p := range points {
		if CanAccess(p, points, 4) {
			accessible = append(accessible, p)
		}
	}

	if len(accessible) == 0 {
		return removed
	}

	// fmt.Printf("Accessible: %d\r\n", len(accessible))

	return CountAccessibleRec(difference(points, accessible), removed+len(accessible))
}

func difference(all []point, remove []point) []point {
	return slices.DeleteFunc(all, func(x point) bool {
		return slices.Contains(remove, x)
	})
}
