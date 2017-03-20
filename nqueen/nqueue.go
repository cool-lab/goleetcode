package nqueen

func solveNQueens(n int) [][]string {

	list := []int{}

	return dfs(list, n)
}

func dfs(list []int, n int) (result [][]string) {
	if len(list) == n {
		result = append(result, (translate(list, n)))
	}
	for col := 0; col < n; col++ {
		if !isValid(list, col) {
			continue
		}
		list = append(list, col)
		result = append(result, dfs(list, n)...)
		list = list[:len(list)-1]
	}
	return
}

func translate(list []int, n int) (ret []string) {

	for _, value := range list {
		b := make([]byte, n)
		for j := 0; j < n; j++ {
			if j != value {
				b[j] = '.'
			} else {
				b[j] = 'Q'
			}
		}

		ret = append(ret, string(b))
	}
	return
}

func isValid(list []int, col int) bool {
	row := len(list)
	for i := 0; i < row; i++ {
		if (list[i] == col) || (i-list[i] == row-col) || (i+list[i] == row+col) {
			return false
		}
	}
	return true
}
