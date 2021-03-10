package main

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
