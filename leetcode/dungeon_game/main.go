package dungeon_game

import (
	"fmt"
	"math"
)

type coord struct {
	x, y int
}

type point struct {
	c    coord
	path []coord
}

func (c coord) String() string {
	return fmt.Sprintf("x: %d; y: %d\n", c.x, c.y)
}

func move(grid [][]int, p point) []point {
	ret := make([]point, 0)
	if p.c.x+1 < len(grid[p.c.y]) {
		c := coord{x: p.c.x + 1, y: p.c.y}
		path := make([]coord, len(p.path))
		copy(path, p.path)
		path = append(path, c)
		ret = append(ret, point{c: c, path: path})
	}
	if p.c.y+1 < len(grid) {
		c := coord{x: p.c.x, y: p.c.y + 1}
		path := make([]coord, len(p.path))
		copy(path, p.path)
		path = append(path, c)
		ret = append(ret, point{c: c, path: path})
	}
	return ret
}

func findAllPath(dungeon [][]int) [][]coord {
	c := coord{0, 0}
	start := point{c: c, path: []coord{c}}
	// If we use seen here, we will not get all the paths...
	//seen := make(map[point]bool)
	path := []point{start}
	goal := coord{y: len(dungeon) - 1, x: len(dungeon[0]) - 1}
	allPaths := make([][]coord, 0)
	for len(path) > 0 {
		var current point
		current, path = path[0], path[1:]
		if current.c == goal {
			allPaths = append(allPaths, current.path)
		}
		paths := move(dungeon, current)
		path = append(path, paths...)
	}
	return allPaths
}

func calculateMinimumHP(dungeon [][]int) int {
	allPaths := findAllPath(dungeon)
	return findMinimumHP(allPaths, dungeon)
}

func findMinimumHP(allPaths [][]coord, dungeon [][]int) int {
	minAge := math.MaxInt64
	for _, p := range allPaths {
		startingHP := 1
		sum := 0
		for _, c := range p {
			sum += dungeon[c.y][c.x]
			if v := startingHP + sum; v <= 0 {
				startingHP += (-v) + 1
			}
		}
		if startingHP < minAge {
			minAge = startingHP
		}
	}
	return minAge
}
