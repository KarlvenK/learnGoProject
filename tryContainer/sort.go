package main

import (
	"fmt"
	"sort"
)

type stuScore struct {
	name  string
	score int
}

type stuScores []stuScore

func (s stuScores) Len() int {
	return len(s)
}

func (s stuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s stuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func trySort() {
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
	fmt.Println("===================")
}
