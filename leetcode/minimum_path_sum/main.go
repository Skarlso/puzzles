package main

import "math"

type point struct {
	x, y int
}

func move(grid [][]int, p point) []point {
	ret := make([]point, 0)
	if p.x+1 < len(grid[p.y]) {
		ret = append(ret, point{x: p.x + 1, y: p.y})
	}
	if p.y+1 < len(grid) {
		ret = append(ret, point{x: p.x, y: p.y + 1})
	}
	return ret
}

func minPathSum(grid [][]int) int {
	start := point{x: 0, y: 0}
	cameFrom := make(map[point]*point)
	cameFrom[start] = nil
	maxY := len(grid)
	maxX := len(grid[maxY-1])
	goal := point{x: maxX - 1, y: maxY - 1}
	path := []point{start}
	min := math.MaxInt64
	for len(path) > 0 {
		var (
			current point
		)

		current, path = path[0], path[1:]
		if current == goal {
			// compare min here if smaller, min is now that
			currentMin := grid[0][0]
			// follow back the breadcrumbs while counting the min
			back := goal
			for back != start {
				currentMin += grid[back.y][back.x]
				back = *cameFrom[back]
			}
			if currentMin < min {
				min = currentMin
			}
			continue
			//break
		}
		paths := move(grid, current)
		for _, next := range paths {
			if _, ok := cameFrom[next]; !ok {
				cameFrom[next] = &current
				path = append(path, next)
			}
		}
	}

	return min
}
