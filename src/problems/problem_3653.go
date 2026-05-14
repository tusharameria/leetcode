// 3653. XOR After Range Multiplication Queries I

package problems

func Problem_3653() {
	nums := []int{2, 3, 1, 5, 4}
	queries := [][]int{
		{1, 4, 2, 3},
		{0, 2, 1, 2},
	}
	println(xorAfterQueriesOld(nums, queries))
}

func xorAfterQueriesOld(nums []int, queries [][]int) int {
	res := 0
	divider := 1000000007
	for i := range queries {
		l, r, step, multiplier := queries[i][0], queries[i][1], queries[i][2], queries[i][3]
		for j := l; j <= r; j += step {
			nums[j] = (nums[j] * multiplier) % divider
		}
	}
	for _, v := range nums {
		res ^= v
	}
	return res
}
