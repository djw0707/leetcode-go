package arrayMultiply

// 从前向后遍历，从后向前遍历计算累积 不允许使用除法
func productExceptSelf(nums []int) []int {
	// 切片固定长度为nums才可以通过下标访问，如果长度为0需要通过append
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left[i] = nums[i]
		} else {
			left[i] = left[i-1] * nums[i]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			right[i] = nums[i]
		} else {
			right[i] = right[i+1] * nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans[i] = right[i+1]
		} else if i == len(nums)-1 {
			ans[i] = left[i-1]
		} else {
			ans[i] = left[i-1] * right[i+1]
		}
	}
	return ans
}
