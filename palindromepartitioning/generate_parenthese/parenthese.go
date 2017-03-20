package generate_parenthese

func generateParenthesis(n int) []string {
	return dfs(n, n, "")
}

func dfs(leftNum, rightNum int, s string) (result []string) {
	if leftNum == 0 && rightNum == 0 {
		result = append(result, s)
	}
	if leftNum > rightNum {
		return
	}
	if leftNum > 0 {
		result = append(result, dfs(leftNum-1, rightNum, s+"(")...)
	}
	if rightNum > 0 {
		result = append(result, dfs(leftNum, rightNum-1, s+")")...)
	}
	return
}
