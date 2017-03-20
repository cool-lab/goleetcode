package subsets2

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)
	return dfs(nums, []int{}, 0)
}

func dfs(nums, list []int, pos int) (result [][]int) {
	copyList := make([]int, len(list))
	copy(copyList, list)
	result = append(result, copyList)

	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] { //当不是第一元素，且是第一个元素之后的相识元素都不放入。
			continue
		}
		list = append(list, nums[i])
		result = append(result, dfs(nums, list, i+1)...)
		list = list[:len(list)-1]
	}

	return result
}
