package permutations

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	list := []int{}
	dfs(&res, list, nums)
	return res
}

func dfs(res *[][]int, list, nums []int) {
	if len(list) == len(nums) {
		replia := make([]int, len(list)) //要进行深度copy.不然会重复。
		copy(replia, list)
		*res = append(*res, replia)
		return
	}
	for i := 0; i < len(nums); i++ {
		if contains(nums[i], list) {
			continue
		}
		list = append(list, nums[i])
		dfs(res, list, nums)
		list = list[0 : len(list)-1]

	}
	return
}

func contains(num int, nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}
