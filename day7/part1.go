package main

import (
	"maps"
	"slices"
)

type coord struct {
	x, y int
}

type grid map[coord]rune

func Part1(lines []string) int {
	start, grid := Parse(lines)
	output := CountSplits(grid, []coord{start}, len(lines), start.y+1, 0)
	return output
}

func CountSplits(g grid, beams []coord, maxy, y int, acc int) int {
	if y >= maxy {
		return acc
	}
	nextBeams := make(map[coord]struct{})

	for _, c := range beams {
		targetCoord := coord{x: c.x, y: y}
		targetVal := g[targetCoord]
		if targetVal == '^' {
			acc++
			nextBeams[coord{x: targetCoord.x - 1, y: y}] = struct{}{}
			nextBeams[coord{x: targetCoord.x + 1, y: y}] = struct{}{}
			continue
		}
		nextBeams[targetCoord] = struct{}{}
	}
	beamCoords := slices.Collect(maps.Keys(nextBeams))
	// fmt.Printf("y=%d beams=%v acc=%d\n", y, beamCoords, acc)

	return CountSplits(g, beamCoords, maxy, y+1, acc)
}

func Parse(lines []string) (coord, grid) {
	height := len(lines)
	width := len(lines[0])

	g := grid{}
	var start coord

	for y := range height {
		for x := range width {
			c := rune(lines[y][x])
			p := coord{x: x, y: y}
			if c == '^' {
				g[p] = c
			}
			if c == 'S' {
				start = p
			}
		}
	}

	return start, g
}
