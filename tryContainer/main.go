package main

import (
	"fmt"
	"sort"
)

func trySet() {
	var set stuScores

	for i := 'a'; i < 'z'; i++ {
		set = append(set, stuScore{string(i), int(i)})
	}
	sort.Sort(set)
	fmt.Println(set)

	sort.Sort(sort.Reverse(set))
	fmt.Println(set)
	sort.Sort(set)
	temp := sort.Search(len(set), func(i int) bool {
		return set[i].score > 101
	})
	fmt.Println(temp)

}
func main() {
	trySet()
}
