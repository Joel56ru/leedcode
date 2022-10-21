package mediumQuestions

import "testing"

type data2 struct {
	result string
	n      int
}

var t4 = []data2{
	{"III", 3},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"I", 1},
	{"V", 5},
	{"IV", 4},
	{"VII", 7},
	{"IX", 9},
	{"XII", 12},
	{"MMMCMXCIX", 3999},
}

func TestIntToRoman(t *testing.T) {
	for i, pair := range t4 {
		k := intToRoman(pair.n)
		if k != pair.result {
			t.Errorf("#%d Not Valid %v. FuncRes: %v", i, pair, k)
		}
	}
}
