package easyQuestions

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i, num := range nums {
		if val, ok := temp[num]; ok {
			return []int{val, i}
		}
		temp[target-num] = i
	}
	return nil
}
