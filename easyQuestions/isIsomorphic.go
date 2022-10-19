package easyQuestions

func isIsomorphic(s string, t string) bool {
	mapST, mapTS := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := mapST[s[i]]; ok {
			if mapST[s[i]] != t[i] {
				return false
			}
		}
		if _, ok := mapTS[t[i]]; ok {
			if mapTS[t[i]] != s[i] {
				return false
			}
		}
		mapST[s[i]], mapTS[t[i]] = t[i], s[i]
	}
	return true
}
