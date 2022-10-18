package mediumQuestions

import (
	"strconv"
)

// https://leetcode.com/problems/count-and-say/submissions/
func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = combine(comparison(res))
	}
	return res
}

type store struct {
	say   rune
	count int
}

func comparison(n string) []store {
	var st []store
	var s rune
	for _, r := range n {
		if s != r {
			st = append(st, store{r, 1})
		} else {
			st[len(st)-1].count++
		}
		s = r
	}
	return st
}

func combine(st []store) string {
	var res string
	for _, s := range st {
		res = res + strconv.Itoa(s.count) + string(s.say)
	}
	return res
}
