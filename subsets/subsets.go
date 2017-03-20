package subsets

func subsets(nums []int) [][]int {
	subset := []int{}
	return dfs(nums, subset, 0)
}

func dfs(nums, subset []int, pos int) (ret [][]int) {

	setCopy := make([]int, len(subset))
	copy(setCopy, subset)
	ret = append(ret, setCopy)

	for i := pos; i < len(nums); i++ {
		subset = append(subset, nums[i])
		ret = append(ret, dfs(nums, subset, i+1)...) //注意返回需要append到内容中
		subset = subset[:len(subset)-1]
	}
	return
}
