package freedomtrail

import (
	"math"
)

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m)
	}

	stack := make([]int, 0)
	for i := 0; i < m; i++ {
		if key[len(key)-1] == ring[i] {
			stack = append(stack, i)
		}
	}

	for i := n - 1; i >= 0; i-- {
		tmp := make([]int, 0)
		for j := 0; j < m; j++ {
			if (i > 0 && key[i-1] == ring[j]) || (i == j && j == 0) {
				dp[i][j] = m * n
				for _, k := range stack {
					dp[i][j] = min(dp[i][j], dp[i+1][k]+min(abs(k-j), m-abs(k-j)))
				}
				tmp = append(tmp, j)
			}
		}
		stack = tmp
	}

	return dp[0][0] + n
}

func min(a ...int) int {
	min := math.MaxInt64
	for _, x := range a {
		if x < min {
			min = x
		}
	}
	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// func dfs(index int, key, ring string) int {
// 	min := 0
// 	sum := 0
// 	rl := len(ring)
// 	i := index
// 	j := index
// 	for ki, r := range key {
// 		var (
// 			steps int
// 		)

// 		// gather all the indexes on the ring for this character
// 		indexes := gatherAllIndex(r, ring)
// 		// fmt.Printf("there are %d indexes for letter %s\n", len(indexes), string(r))
// 		// TODO: Fix keeping track of min... this will be lower since I'm only doing a partial key lookup.
// 		if len(indexes) > 1 {
// 			for _, index := range indexes {
// 				s := dfs(index, key[ki+1:], ring)
// 				if s < min {
// 					min = s
// 				}
// 			}
// 		}

// 		for {
// 			if rune(ring[i]) == r {
// 				j = i
// 				break
// 			}
// 			if rune(ring[j]) == r {
// 				i = j
// 				break
// 			}
// 			i = (i + 1) % rl
// 			j = negativeMod(j-1, rl)
// 			steps++
// 		}
// 		sum += steps
// 	}
// 	return 0
// }

// func gatherAllIndex(r rune, ring string) (indexes []int) {
// 	for i, rr := range ring {
// 		if rr == r {
// 			indexes = append(indexes, i)
// 		}
// 	}
// 	return
// }

// func negativeMod(d, m int) int {
// 	var res int = d % m
// 	if (res < 0 && m > 0) || (res > 0 && m < 0) {
// 		return res + m
// 	}
// 	return res
// }
