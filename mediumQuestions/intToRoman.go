package mediumQuestions

// https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	nM := []nMap{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}
	var res string
	for _, n := range nM {
		for num >= n.num {
			num -= n.num
			res += n.roman
		}
	}
	return res
}

type nMap struct {
	num   int
	roman string
}
