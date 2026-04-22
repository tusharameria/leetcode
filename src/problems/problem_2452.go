// 2452. Words Within Two Edits of Dictionary

package problems

import "fmt"

func Problem_2452() {
	queries := []string{"word", "note", "ants", "wood"}
	dictionary := []string{"wood", "joke", "moat"}
	fmt.Println(twoEditWords(queries, dictionary))
}

func twoEditWords(queries []string, dictionary []string) []string {
	var result []string
	for _, query := range queries {
		for _, word := range dictionary {
			diff := 0
			for i := 0; i <= len(query)-1; i++ {
				if query[i] != word[i] {
					diff++
				}
			}
			if diff <= 2 {
				result = append(result, query)
				break
			}
		}
	}
	return result
}
