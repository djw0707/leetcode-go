package findFirstMissInt

// 将每个数字i放到i-1位置，调整之后如果数字和下标对应关系不一致就是缺少的正数
func firstMissingPositive(nums []int) int {
	ans := -1
	max := 0
	for i := 0; i < len(nums); i++ {
		// 考虑到下标需要保证在数组范围内 且需要调整 nums[i] != nums[nums[i]-1]保证不会死循环
		for nums[i] < len(nums) && nums[i] > 0 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = i + 1
			return ans
		} else {
			max = i + 1
		}
	}
	if ans == -1 {
		ans = max + 1
	}
	return ans
}
