package mediumQuestions

import "testing"

type data struct {
	result string
	n      int
}

var t3 = []data{
	{"1", 1},
	{"11", 2},
	{"21", 3},
	{"1211", 4},
	{"111221", 5},
	{"312211", 6},
	{"13112221", 7},
	{"1113213211", 8},
	{"31131211131221", 9},
	{"13211311123113112211", 10},
}

func TestCheckIfPangram(t *testing.T) {
	for i, pair := range t3 {
		k := countAndSay(pair.n)
		if k != pair.result {
			t.Errorf("#%d Not Valid %v. FuncRes: %v", i, pair, k)
		}
	}
}
