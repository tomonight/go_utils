package basic

func getPatternTable(s string) []int {

	next := make([]int, len(s))
	for i, j := 1, 0; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}

	return next
}

func kmp(s1, s2 string) int {
	next := getPatternTable(s2)
	for i, j := 0, 0; i < len(s1); i++ {

		for j > 0 && s1[i] != s2[j] {
			j = next[j-1]
		}
		if s1[i] == s2[j] {
			j++
		}
		if j == len(s2) {
			return i - j + 1
		}
	}
	return -1
}
