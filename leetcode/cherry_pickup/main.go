package cherry_pickup

import "fmt"

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

func moveDown(ly, lx int, p point) []point {
	ret := make([]point, 0)
	if p.c.x+1 < lx {
		c := coord{x: p.c.x + 1, y: p.c.y}
		ret = append(ret, point{c: c, moveSequence: p.moveSequence + "R"})
	}
	if p.c.y+1 < ly {
		c := coord{x: p.c.x, y: p.c.y + 1}
		ret = append(ret, point{c: c, moveSequence: p.moveSequence + "D"})
	}
	return ret
}

func moveUp(p point) []point {
	ret := make([]point, 0)
	if p.c.x-1 >= 0 {
		c := coord{x: p.c.x - 1, y: p.c.y}
		ret = append(ret, point{c: c, moveSequence: p.moveSequence + "L"})
	}
	if p.c.y-1 >= 0 {
		c := coord{x: p.c.x, y: p.c.y - 1}
		ret = append(ret, point{c: c, moveSequence: p.moveSequence + "U"})
	}
	return ret
}

func cherryPickup(grid [][]int) []string {
	c := coord{0, 0}
	start := point{c: c, moveSequence: "S"}
	path := []point{start}
	goal := coord{y: len(grid) - 1, x: len(grid[0]) - 1}
	allBytePath := make([]string, 0)
	for len(path) > 0 {
		var current point
		current, path = path[0], path[1:]
		if current.c == goal {
			allBytePath = append(allBytePath, current.moveSequence)
		}
		paths := move(grid, current)
		path = append(path, paths...)
	}
	return allBytePath
}
