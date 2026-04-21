//

package problems

import "fmt"

func Problem_1722() {
	source := []int{1, 2, 3}
	target := []int{2, 3, 1}
	allowedSwaps := [][]int{{0, 1}}
	fmt.Println(minimumHammingDistance(source, target, allowedSwaps))
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)

	// --- DSU ---
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) {
		px, py := find(x), find(y)
		if px != py {
			parent[py] = px
		}
	}

	// --- Build components ---
	for _, swap := range allowedSwaps {
		union(swap[0], swap[1])
	}

	// --- Group indices by root ---
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := find(i)
		groups[root] = append(groups[root], i)
	}

	// --- Count mismatches ---
	res := 0

	for _, indices := range groups {
		freq := make(map[int]int)

		// count source values
		for _, idx := range indices {
			freq[source[idx]]++
		}

		// match target values
		for _, idx := range indices {
			if freq[target[idx]] > 0 {
				freq[target[idx]]--
			} else {
				res++
			}
		}
	}

	return res
}
