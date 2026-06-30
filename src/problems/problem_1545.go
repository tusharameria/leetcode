// 1545. Find Kth Bit in Nth Binary String

package problems

import "fmt"

func Problem_1545() {
	n := 4
	k := 14
	fmt.Println(string([]byte{findKthBit(n, k)}))
}

func findKthBit(n int, k int) byte {
	idx := k - 1
	totalLen := pow2(uint32(n)) - 1
	flip := false
	for {
		if idx == 0 {
			if flip {
				return '1'
			}
			return '0'
		}
		mid := totalLen / 2
		if idx == totalLen-1 || idx == mid {
			if flip {
				return '0'
			}
			return '1'
		}
		if idx > mid {
			idx = totalLen - 1 - idx
			flip = !flip
		}
		totalLen = mid
	}
}
