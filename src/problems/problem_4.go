package problems

import "fmt"

func Problem_4() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0
	m := len(nums1)
	n := len(nums2)
	total := m + n
	if total == 0 {
		return 0
	}
	isEven := total%2 == 0
	if m == 0 {
		if !isEven {
			return float64(nums2[(n-1)/2])
		} else {
			return float64(nums2[n/2]+nums2[(n/2)-1]) / 2
		}
	}
	if n == 0 {
		if !isEven {
			return float64(nums1[(m-1)/2])
		} else {
			return float64(nums1[m/2]+nums1[(m/2)-1]) / 2
		}
	}
	new := []int{}
	for i <= m-1 && j <= n-1 {
		if nums1[i] > nums2[j] {
			new = append(new, nums2[j])
			j++
		} else {
			new = append(new, nums1[i])
			i++
		}
		if !isEven {
			if len(new) == (total+1)/2 {
				return float64(new[len(new)-1])
			}
		} else {
			if len(new) == (total/2)+1 {
				return float64((new[len(new)-1] + new[len(new)-2])) / 2
			}
		}
	}
	if i == m {
		new = append(new, nums2[j:]...)
	} else {
		new = append(new, nums1[i:]...)
	}
	if !isEven {
		return float64(new[(total-1)/2])
	} else {
		return float64(new[(total/2)]+new[(total/2)-1]) / 2
	}
}

func findMedianSortedArraysEff(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	total := m + n
	isEven := total%2 == 0
	targetIdx := int(total / 2)

	if m == 0 {
		if !isEven {
			return float64(nums2[targetIdx])
		} else {
			return float64(nums2[targetIdx-1]+nums2[targetIdx]) / 2
		}
	}

	if n == 0 {
		if !isEven {
			return float64(nums1[targetIdx])
		} else {
			return float64(nums1[targetIdx-1]+nums1[targetIdx]) / 2
		}
	}

	i := 0
	j := 0
	lastMin := 0

	for i < m && j < n {
		minVal := min(nums1[i], nums2[j])
		if i+j == targetIdx {
			if isEven {
				return float64(lastMin+minVal) / 2
			} else {
				return float64(minVal)
			}
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
		lastMin = minVal
	}

	if i == m {
		if i+j == targetIdx {
			if isEven {
				return float64(lastMin+nums2[j]) / 2
			} else {
				return float64(nums2[j])
			}
		}
		idx := targetIdx - i
		if !isEven {
			return float64(nums2[idx])
		} else {
			return float64(nums2[idx-1]+nums2[idx]) / 2
		}
	}

	if j == n {
		if i+j == targetIdx {
			if isEven {
				return float64(lastMin+nums1[i]) / 2
			} else {
				return float64(nums1[i])
			}
		}
		idx := targetIdx - j
		if !isEven {
			return float64(nums1[idx])
		} else {
			return float64(nums1[idx-1]+nums1[idx]) / 2
		}
	}

	return 0
}
