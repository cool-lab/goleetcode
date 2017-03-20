package palindromepartitioning

func partition(s string) [][]string {
	sub := []string{}
	return subsets(s, sub, 0)
}

func subsets(s string, sub []string, pos int) (result [][]string) {
	if pos == len(s) {
		copySub := make([]string, len(sub))
		copy(copySub, sub)
		result = append(result, copySub)
		return
	}

	for i := pos; i < len(s); i++ {
		subString := s[pos : i+1]
		if !isPalindrome(subString) {
			continue
		}
		sub = append(sub, subString)
		result = append(result, subsets(s, sub, pos+1)...)
		sub = sub[:len(sub)-1]
	}

	return
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
