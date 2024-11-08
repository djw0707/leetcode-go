package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	cGraph := make([][]int, numCourses)
	degree := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		cGraph[i] = make([]int, 0)
	}
	for _, pre := range prerequisites {
		cGraph[pre[1]] = append(cGraph[pre[1]], pre[0])
		degree[pre[0]]++
	}
	flag := make([]bool, numCourses)
	ans := make([]int, 0)
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		top := queue[0]
		ans = append(ans, top)
		queue = queue[1:]
		for _, index := range cGraph[top] {
			if flag[index] == true {
				continue
			}
			flag[index] = true
			queue = append(queue, index)
		}
	}
	return ans
}

func main() {
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(4, prerequisites))
}
