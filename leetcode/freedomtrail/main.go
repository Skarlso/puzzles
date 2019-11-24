package freedomtrail

import (
	"math"
)

// Integer[][] dp = null;
// public int findRotateSteps(String ring, String key) {
// 	dp = new Integer[ring.length()][key.length()];
// 	return steps(0, ring, 0, key);
// }
// private int steps(int i, String ring, int j, String key){
// 	if(j == key.length())
// 		return 0;
// 	if(dp[i][j] != null) return dp[i][j];
// 	dp[i][j] = Integer.MAX_VALUE;
// 	char ch = key.charAt(j);
// 	for(int k=0; k < ring.length(); k++){
// 		int tmp = (k+i) % ring.length();
// 		if(ring.charAt(tmp) == ch)
// 			dp[i][j] = Math.min(dp[i][j], 1 + Math.min(k, ring.length()-k) + steps(tmp, ring, j+1, key));
// 	}
// 	return dp[i][j];
// }
var dp [][]int

func findRotateSteps(ring string, key string) int {
	dp = make([][]int, len(ring))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(key))
	}

	return sum(0, 0, ring, key)
}

func sum(i, j int, ring, key string) int {
	if j == len(key) {
		return 0
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	dp[i][j] = math.MaxInt64
	c := key[j]
	for k := 0; k < len(ring); k++ {
		tmp := (k + i) % len(ring)
		if ring[tmp] == c {
			dp[i][j] = min(dp[i][j], 1+min(k, len(ring)-k)+sum(tmp, j+1, ring, key))
		}
	}

	return dp[i][j]
}

// func findRotateSteps(ring string, key string) int {
// 	m, n := len(ring), len(key)
// 	dp := make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		dp[i] = make([]int, m)
// 	}

// 	stack := make([]int, 0)
// 	for i := 0; i < m; i++ {
// 		if key[len(key)-1] == ring[i] {
// 			stack = append(stack, i)
// 		}
// 	}

// 	for i := n - 1; i >= 0; i-- {
// 		tmp := make([]int, 0)
// 		for j := 0; j < m; j++ {
// 			if (i > 0 && key[i-1] == ring[j]) || (i == j && j == 0) {
// 				dp[i][j] = m * n
// 				for _, k := range stack {
// 					dp[i][j] = min(dp[i][j], dp[i+1][k]+min(abs(k-j), m-abs(k-j)))
// 				}
// 				tmp = append(tmp, j)
// 			}
// 		}
// 		stack = tmp
// 	}

// 	return dp[0][0] + n
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
