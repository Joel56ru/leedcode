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
				break
			}
		}
	}
}

type node struct {
	node   *ListNode
	result *ListNode
}

// {&ListNode{1, nil}
var t2 = []node{
	{&ListNode{1, &ListNode{1, &ListNode{2, nil}}}, &ListNode{1, &ListNode{2, nil}}},
}

func TestDeleteDuplicates(t *testing.T) {
	for i, pair := range t2 {
		nodeRes := deleteDuplicates(pair.node)
		for nodeRes != nil || pair.result != nil {
			if nodeRes.Val != pair.result.Val {
				t.Errorf("#%d Not Valid %v. FuncRes: %v", i, nodeRes.Val, pair.result.Val)
			}
			nodeRes = nodeRes.Next
			pair.result = pair.result.Next
		}
	}
}

type stringBool struct {
	s      string
	result bool
}

var t3 = []stringBool{
	{"thequickbrownfoxjumpsoverthelazydog", true},
	{"leetcode", false},
	{"qwertyuiopasdfghjklzxcvbnm", true},
}

func TestCheckIfPangram(t *testing.T) {
	for i, pair := range t3 {
		k := checkIfPangram(pair.s)
		if k != pair.result {
			t.Errorf("#%d Not Valid %v. FuncRes: %v", i, pair, k)
		}
	}
}
