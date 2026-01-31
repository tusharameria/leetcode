package problems

func Problem_766() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')
	println(nextGreatestLetter(letters, target))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	len := len(letters)
	if target < letters[0] || target >= letters[len-1] {
		return letters[0]
	}
	left, right := 0, len-1
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return letters[left]
}
