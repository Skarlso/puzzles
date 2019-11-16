package dungeon_game

import (
	"container/heap"
	"fmt"
)

type coord struct {
	x, y int
}

func (c coord) String() string {
	return fmt.Sprintf("x: %d; y: %d\n", c.x, c.y)
}

type point struct {
	c        coord
	index    int
	priority int
}

type ppQ []point

func (pq ppQ) Len() int { return len(pq) }
func (pq ppQ) Less(i, j int) bool {
	// implement the cost logic here?
	return pq[i].priority < pq[j].priority
}

func (pq ppQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *ppQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(point)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *ppQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func move(grid [][]int, p point) []point {
	ret := make([]point, 0)
	if p.c.x+1 < len(grid[p.c.y]) {
		ret = append(ret, point{c: coord{x: p.c.x + 1, y: p.c.y}})
	}
	if p.c.y+1 < len(grid) {
		ret = append(ret, point{c: coord{x: p.c.x, y: p.c.y + 1}})
	}
	return ret
}

func calculateMinimumHP(dungeon [][]int) int {
	// find the path that leads to the smallest positive integer ( absolute value smallest ? )
	// and go through that until the result is 1.
	// put this whole thing into a loop?
	knightHP := 0
	start := point{c: coord{x: 0, y: 0}, priority: 0, index: 0}
	cameFrom := map[coord]coord{
		start.c: start.c,
	}
	costSoFar := map[coord]int{
		start.c: 0,
	}
	knightHP += dungeon[start.c.y][start.c.x]
	maxY := len(dungeon)
	maxX := len(dungeon[maxY-1])
	goal := coord{x: maxX - 1, y: maxY - 1}
	path := make(ppQ, 1)
	path[0] = start
	heap.Init(&path)
	for path.Len() > 0 {
		current := path.Pop().(point)
		knightHP += dungeon[current.c.y][current.c.x]
		if current.c == goal {
			fmt.Printf("cost so far: %d\n", costSoFar[current.c])
			continue
			//break
		}
		// move needs to add in paths so it is tracked what route has been taken so far.
		paths := move(dungeon, current)
		for _, next := range paths {
			newCost := costSoFar[current.c] + dungeon[next.c.y][next.c.x]
			if _, ok := cameFrom[next.c]; !ok || newCost < costSoFar[next.c] {
				cameFrom[next.c] = current.c
				next.priority = newCost
				costSoFar[next.c] = newCost
				path.Push(next)
			}
		}
	}

	p := make([]coord, 0)
	current := goal
	sum := 0
	for current != start.c {
		p = append(p, current)
		sum += dungeon[current.y][current.x]
		current = cameFrom[current]
	}
	p = append(p, start.c)
	//sum += dungeon[start.c.y][start.c.x]
	fmt.Println(p, sum)
	return 0
}
