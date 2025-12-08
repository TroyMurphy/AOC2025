package main

func Part2(lines []string) int {
	start, grid := Parse(lines)
	output := CountWorlds(grid, start, len(lines), start.y+1)
	return output
}

var solutions map[coord]int = make(map[coord]int)

func CountWorlds(g grid, point coord, maxy, y int) int {
	if y >= maxy {
		return 1
	}

	targetCoord := coord{x: point.x, y: y}
	targetVal := g[targetCoord]

	if solutions[targetCoord] > 0 {
		return solutions[targetCoord]
	}

	if targetVal == '^' {
		value := CountWorlds(g, coord{point.x - 1, y}, maxy, y+1) +
			CountWorlds(g, coord{x: point.x + 1, y: y}, maxy, y+1)
		solutions[targetCoord] = value
		return value
	}
	value := CountWorlds(g, coord{x: point.x, y: y}, maxy, y+1)
	solutions[targetCoord] = value
	return value
}
