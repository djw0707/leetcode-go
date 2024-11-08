package mergeIntervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return false
		}
	})
	ans := make([][]int, 0)
	// 在这里先把第一个加进来避免漏掉最后一个
	ans = append(ans, intervals[0])
	for _, interval := range intervals {
		start, end := ans[len(ans)-1][0], ans[len(ans)-1][1]
		if interval[0] <= end && interval[1] > end {
			end = interval[1]
			ans[len(ans)-1] = []int{start, end}
		} else if interval[0] > end {
			ans = append(ans, interval)
		}
	}
	return ans
}
