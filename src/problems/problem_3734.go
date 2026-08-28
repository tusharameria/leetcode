// 3734. Lexicographically Smallest Palindromic Permutation Greater Than Target

package problems

func lexPalindromicPermutation(s string, target string) string {
	frequency := make([]int, 26)
	for i := 0; i < len(s); i++ {
		frequency[s[i]-'a']++
	}

	middle := byte(0)
	oddCount := 0
	for c := 0; c < 26; c++ {
		if frequency[c]%2 == 1 {
			oddCount++
			middle = byte('a' + c)
		}
	}

	if oddCount > 1 {
		return ""
	}

	halfCount := make([]int, 26)
	for c := 0; c < 26; c++ {
		halfCount[c] = frequency[c] / 2
	}

	k := len(s) / 2
	targetHalf := target[:k]

	smallestGreaterOrEqual := func(originalCount []int, prefix string) string {
		count := make([]int, 26)
		copy(count, originalCount)

		matched := 0

		for matched < k && count[prefix[matched]-'a'] > 0 {
			count[prefix[matched]-'a']--
			matched++
		}

		if matched == k {
			return prefix
		}

		for pos := matched; pos >= 0; pos-- {
			if pos < matched {
				count[prefix[pos]-'a']++
			}

			for c := int(prefix[pos]-'a') + 1; c < 26; c++ {
				if count[c] == 0 {
					continue
				}

				result := make([]byte, 0, k)
				result = append(result, prefix[:pos]...)
				result = append(result, byte('a'+c))
				count[c]--

				for ch := 0; ch < 26; ch++ {
					for times := 0; times < count[ch]; times++ {
						result = append(result, byte('a'+ch))
					}
				}

				return string(result)
			}
		}

		return ""
	}

	buildPalindrome := func(half string) string {
		result := make([]byte, 0, len(s))
		result = append(result, half...)

		if middle != 0 {
			result = append(result, middle)
		}

		for i := len(half) - 1; i >= 0; i-- {
			result = append(result, half[i])
		}

		return string(result)
	}

	nextPermutation := func(chars []byte) bool {
		pivot := len(chars) - 2

		for pivot >= 0 && chars[pivot] >= chars[pivot+1] {
			pivot--
		}

		if pivot < 0 {
			return false
		}

		swapPos := len(chars) - 1

		for chars[swapPos] <= chars[pivot] {
			swapPos--
		}

		chars[pivot], chars[swapPos] = chars[swapPos], chars[pivot]

		left, right := pivot+1, len(chars)-1
		for left < right {
			chars[left], chars[right] = chars[right], chars[left]
			left++
			right--
		}

		return true
	}

	half := smallestGreaterOrEqual(halfCount, targetHalf)

	if half == "" && k > 0 {
		return ""
	}

	candidate := buildPalindrome(half)

	if candidate > target {
		return candidate
	}

	chars := []byte(half)
	if !nextPermutation(chars) {
		return ""
	}

	return buildPalindrome(string(chars))
}
