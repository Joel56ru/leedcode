package easyQuestions

// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
func checkIfPangram(sentence string) bool {
	l := make(map[int32]struct{})
	for _, s := range sentence {
		if _, ok := l[s]; !ok {
			l[s] = struct{}{}
		}
	}
	return len(l) == 26
}
