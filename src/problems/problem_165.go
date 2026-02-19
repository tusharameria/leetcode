package problems

import (
	"fmt"
)

// 165. Compare Version Numbers

func Problem_165() {
	version1 := "1.20"
	version2 := "1.20"
	fmt.Println("res : ", compareVersion(version1, version2))
}

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	sum1 := 0
	sum2 := 0

	for {
		if i > len(version1)-1 && j > len(version2)-1 {
			if sum1 != sum2 {
				if sum1 < sum2 {
					return -1
				}
				return 1
			} else {
				return 0
			}
		}
		if (i <= len(version1)-1) != (j <= len(version2)-1) {
			if (i <= len(version1)-1 && version1[i] == '.') || (j <= len(version2)-1 && version2[j] == '.') {
				if sum1 != sum2 {
					if sum1 < sum2 {
						return -1
					}
					return 1
				}
				sum1 = 0
				sum2 = 0
				i++
				j++
			}
		}
		if i <= len(version1)-1 && j <= len(version2)-1 && version1[i] == '.' && version2[j] == '.' {
			if sum1 != sum2 {
				if sum1 < sum2 {
					return -1
				}
				return 1
			}
			sum1, sum2 = 0, 0
			i++
			j++
		}
		if i <= len(version1)-1 && version1[i] != '.' {
			sum1 = sum1*10 + int(version1[i]-'0')
			i++
		}
		if j <= len(version2)-1 && version2[j] != '.' {
			sum2 = sum2*10 + int(version2[j]-'0')
			j++
		}
	}
}
