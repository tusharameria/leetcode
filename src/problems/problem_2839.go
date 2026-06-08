// 2839. Check if Strings Can be Made Equal With Operations I

package problems

import "fmt"

func Problem_2839() {
	s1 := "abcd"
	s2 := "dacb"
	fmt.Println(canBeEqual(s1, s2))
}

func canBeEqual(s1 string, s2 string) bool {
	oddPlaces := make([]int, 26)
	evenPlaces := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		alpha := int(s1[i] - 'a')
		if i%2 == 0 {
			evenPlaces[alpha]++
		} else {
			oddPlaces[alpha]++
		}
	}

	for i := 0; i < len(s2); i++ {
		alpha := int(s2[i] - 'a')
		if i%2 == 0 {
			if evenPlaces[alpha] == 0 {
				return false
			}
			evenPlaces[alpha]--
		} else {
			if oddPlaces[alpha] == 0 {
				return false
			}
			oddPlaces[alpha]--
		}
	}

	return true
}
