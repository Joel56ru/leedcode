package easyQuestions

import "testing"

type data struct {
	s      string
	t      string
	result bool
}

var t4 = []data{
	{"egg", "add", true},
	{"foo", "bar", false},
	{"paper", "title", true},
	{"a", "b", true},
}

func TestIsIsomorphic(t *testing.T) {
	for i, pair := range t4 {
		k := isIsomorphic(pair.s, pair.t)
		if k != pair.result {
			t.Errorf("#%d Not Valid %v. FuncRes: %v", i, pair, k)
		}
	}
}
