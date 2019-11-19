package cherry_pickup

import (
	"container/heap"
)

type coordPQ struct {
	x, y int
}

type pointPQ struct {
	c        coordPQ
	index    int
	priority int
}

type ppQ []pointPQ

// Len returns length
func (pq ppQ) Len() int { return len(pq) }

// Less returns less. This is reversed, because I want it to move towards
// the 1s.
func (pq ppQ) Less(i, j int) bool {
	// implement the cost logic here?
	return pq[i].priority > pq[j].priority
}

// Swap swaps elements
func (pq ppQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push pushes an item based on priority
func (pq *ppQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(pointPQ)
	item.index = n
	*pq = append(*pq, item)
}

// Pop pops an item based on priority
func (pq *ppQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func moveDown(grid [][]int, p pointPQ) []pointPQ {
	ret := make([]pointPQ, 0)
	if p.c.x+1 < len(grid[p.c.y]) && grid[p.c.y][p.c.x+1] > -1 {
		ret = append(ret, pointPQ{c: coordPQ{x: p.c.x + 1, y: p.c.y}})
	}
	if p.c.y+1 < len(grid) && grid[p.c.y+1][p.c.x] > -1 {
		ret = append(ret, pointPQ{c: coordPQ{x: p.c.x, y: p.c.y + 1}})
	}
	return ret
}

func moveUp(grid [][]int, p pointPQ) []pointPQ {
	ret := make([]pointPQ, 0)
	if p.c.x-1 >= 0 && grid[p.c.y][p.c.x-1] > -1 {
		ret = append(ret, pointPQ{c: coordPQ{x: p.c.x - 1, y: p.c.y}})
	}
	if p.c.y-1 >= 0 && grid[p.c.y-1][p.c.x] > -1 {
		ret = append(ret, pointPQ{c: coordPQ{x: p.c.x, y: p.c.y - 1}})
	}
	return ret
}

func cherryPickup(grid [][]int) int {
	// Travers top down
	start := coordPQ{x: 0, y: 0}
	goal := coordPQ{x: len(grid[0]) - 1, y: len(grid) - 1}
	sum := traversMaze(moveDown, pointPQ{c: start, priority: 0, index: 0}, goal, grid)
	sum += traversMaze(moveUp, pointPQ{c: goal, priority: 0, index: 0}, start, grid)
	return sum
}

func traversMaze(mover func([][]int, pointPQ) []pointPQ, start pointPQ, goal coordPQ, grid [][]int) int {
	cameFrom := map[coordPQ]coordPQ{
		start.c: start.c,
	}
	costSoFar := map[coordPQ]int{
		start.c: 0,
	}
	path := make(ppQ, 1)
	path[0] = start
	heap.Init(&path)
	for path.Len() > 0 {
		current := path.Pop().(pointPQ)
		if current.c == goal {
			continue
		}
		// move needs to add in paths so it is tracked what route has been taken so far.
		paths := mover(grid, current)
		if current.c != goal && len(paths) < 1 {
			return 0
		}
		for _, next := range paths {
			newCost := costSoFar[current.c] + grid[next.c.y][next.c.x]
			if _, ok := cameFrom[next.c]; !ok || newCost > costSoFar[next.c] {
				cameFrom[next.c] = current.c
				next.priority = newCost
				costSoFar[next.c] = newCost
				path.Push(next)
			}
		}
	}

	current := goal
	if current == start.c {
		if sum := grid[current.y][current.x]; sum > 0 {
			grid[current.y][current.x] = 0
			return sum
		}
	}
	var sum int
	for current != start.c {
		sum += grid[current.y][current.x]
		grid[current.y][current.x] = 0
		current = cameFrom[current]
	}
	return sum
}
