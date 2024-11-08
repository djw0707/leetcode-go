package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 如果m+n是奇数 找第(m+n)/2大的数
	// 如果m+n是偶数 找第(m+n)/2和第(m+n)/2+1大的数
	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		return float64(finKMaxNum(nums1, nums2, l/2+1))
	} else {
		return float64(finKMaxNum(nums1, nums2, l/2)+finKMaxNum(nums1, nums2, l/2+1)) / 2
	}
	return 0
}

func finKMaxNum(nums1 []int, nums2 []int, k int) int {
	// 两个指针记录当前查找位置
	ptr1, ptr2 := 0, 0
	for {
		// 如果有一个数组已经查找到结尾，直接去另外一个数组查找
		if ptr1 == len(nums1) {
			return nums2[ptr2+k-1]
		}
		if ptr2 == len(nums2) {
			return nums1[ptr1+k-1]
		}
		// 如果k等于1说明要找第1小的数字，就在当前两个指针中比较返回小的
		if k == 1 {
			return min(nums1[ptr1], nums2[ptr2])
		}
		// 根据大小来选择移动指针，移除第k-1小前面的数字
		newptr1, newptr2 := min(ptr1+k/2, len(nums1))-1, min(ptr2+k/2, len(nums2))-1
		if nums1[newptr1] <= nums2[newptr2] {
			k = k - (newptr1 - ptr1 + 1)
			ptr1 = newptr1 + 1
		} else if nums1[newptr1] > nums2[newptr2] {
			k = k - (newptr2 - ptr2 + 1)
			ptr2 = newptr2 + 1
		}
	}
	return 0
}



func max(a int, b int) int{
    if a>b {
        return a
    }
    return b
}

func maxProfit(prices []int) int {
    // dp[i][j][k] 表示第n天完成j次交易的最大利润,此时手里是否有股票为k可以取0或1
    dp:=make([][3][2]int, len(prices)+1)
    for i:=0;i<=len(prices);i++{
        for j:=1;j<=2;j++{
            if i==0{
                dp[i][j][0] = 0
                dp[i][j][1] = -prices[i]
                continue
            }
            dp[i][j][1] = max(dp[i-1][j][1],dp[i-1][j-1][0]-prices[i-1])
            dp[i][j][0] = max(dp[i-1][j][1]+prices[i-1],dp[i-1][j][0])
        }
    }
    return dp[len(prices)][2][0]
}