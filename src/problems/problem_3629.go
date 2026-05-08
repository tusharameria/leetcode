// 3629. Minimum Jumps to Reach End via Prime Teleportation

package problems

import "fmt"

func Problem_3629() {
	nums := []int{10}
	fmt.Println(minJumps(nums))
}

func minJumps(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n - 1
	}
	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	factorIndices := make(map[int][]int)
	spf := buildSPF(max)
	for i, val := range nums {
		factors := getFactors(val, spf)
		for _, factor := range factors {
			factorIndices[factor] = append(factorIndices[factor], i)
		}
	}
	visited := make([]bool, n)
	queue := []int{0}
	visited[0] = true

	steps := 0

	for len(queue) > 0 {
		size := len(queue)

		for s := 0; s < size; s++ {
			i := queue[0]
			queue = queue[1:]

			if i == n-1 {
				return steps
			}

			// i-1
			if i-1 >= 0 && !visited[i-1] {
				visited[i-1] = true
				queue = append(queue, i-1)
			}

			// i+1
			if i+1 < n && !visited[i+1] {
				visited[i+1] = true
				queue = append(queue, i+1)
			}

			if isPrime(nums[i]) {
				// teleportation
				factors := getFactors(nums[i], spf)

				for _, f := range factors {
					for _, nxt := range factorIndices[f] {
						if !visited[nxt] {
							visited[nxt] = true
							queue = append(queue, nxt)
						}
					}

					// IMPORTANT optimization
					delete(factorIndices, f)
				}
			}
		}

		steps++
	}
	return steps
}

func buildSPF(num int) []int {
	spf := make([]int, num+1)
	for i := 2; i <= num; i++ {
		if spf[i] == 0 {
			spf[i] = i
			for j := i * i; j <= num; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}
	return spf
}

func getFactors(num int, spf []int) []int {
	res := []int{}
	for num > 1 {
		p := spf[num]
		res = append(res, p)
		for num%p == 0 {
			num /= p
		}
	}
	return res
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
