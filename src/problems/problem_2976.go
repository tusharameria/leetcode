package problems

import "fmt"

// 2976. Minimum Cost to Convert String I

func Problem_2976() {
	source := "abcd"
	target := "acbe"
	original := []byte{'a', 'b', 'c', 'c', 'e', 'd'}
	changed := []byte{'b', 'c', 'b', 'e', 'b', 'e'}
	cost := []int{2, 5, 5, 1, 2, 20}
	// original := []byte{'a', 'b', 'c', 'c', 'e', 'd'}
	// changed := []byte{'b', 'c', 'b', 'e', 'b', 'e'}
	// cost := []int{2, 5, 5, 1, 2, 20}
	fmt.Printf("minimumCost : %v\n", minimumCost(source, target, original, changed, cost))
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	const INF int64 = 1e18
	n := 26

	// dist[i][j] = min cost to convert char i -> j
	dist := make([][]int64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}

	// direct transformations
	for i := 0; i < len(original); i++ {
		u := int(original[i] - 'a')
		v := int(changed[i] - 'a')
		if int64(cost[i]) < dist[u][v] {
			dist[u][v] = int64(cost[i])
		}
	}

	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	var ans int64 = 0
	for i := 0; i < len(source); i++ {
		s := int(source[i] - 'a')
		t := int(target[i] - 'a')
		if dist[s][t] == INF {
			return -1
		}
		ans += dist[s][t]
	}

	return ans
}
