package easyQuestions

import "testing"

type intsIntInt struct {
	nums   []int
	target int
	result []int
}

var t1 = []intsIntInt{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for i, pair := range t1 {
		k := twoSum(pair.nums, pair.target)
		for j, num := range pair.result {
			if num != k[j] {
				t.Errorf("#%d Not Valid %v. FuncRes: %v", i, pair, k)
				continue
			}
		}
	}
}
