package rotateArray

func reverseArr(nums []int, i int, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	//翻转三次数组
	i := k % len(nums)
	reverseArr(nums, 0, len(nums)-1)
	reverseArr(nums, 0, i-1)
	reverseArr(nums, i, len(nums)-1)
}
