package dungeon_game

import (
	"fmt"
	"math"
)

type coord struct {
	x, y int
}

type point struct {
	c            coord
	moveSequence string
}

func (c coord) String() string {
	return fmt.Sprintf("x: %d; y: %d\n", c.x, c.y)
}

func move(grid [][]int, p point) []point {
	ret := make([]point, 0)
	if p.c.x+1 < len(grid[p.c.y]) {
		c := coord{x: p.c.x + 1, y: p.c.y}
		ret = append(ret, point{c: c, moveSequence: p.moveSequence + "R"})
	}
	if p.c.y+1 < len(grid) {
		c := coord{x: p.c.x, y: p.c.y + 1}
		ret = append(ret, point{c: c, moveSequence: p.moveSequence + "D"})
	}
	return ret
}

func findAllPath(dungeon [][]int) []string {
	c := coord{0, 0}
	start := point{c: c, moveSequence: "S"}
	path := []point{start}
	goal := coord{y: len(dungeon) - 1, x: len(dungeon[0]) - 1}
	allBytePath := make([]string, 0)
	for len(path) > 0 {
		var current point
		current, path = path[0], path[1:]
		if current.c == goal {
			allBytePath = append(allBytePath, current.moveSequence)
		}
		paths := move(dungeon, current)
		path = append(path, paths...)
	}
	return allBytePath
}

func calculateMinimumHP(dungeon [][]int) int {
	allPaths := findAllPath(dungeon)
	return findMinimumHP(allPaths, dungeon)
}

func findMinimumHP(allPaths []string, dungeon [][]int) int {
	minAge := math.MaxInt64
	for _, p := range allPaths {
		startingHP := 1
		var (
			x, y, sum int
		)
		for _, c := range p {
			switch c {
			case 'R':
				x++
			case 'D':
				y++
			}
			sum += dungeon[y][x]
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
