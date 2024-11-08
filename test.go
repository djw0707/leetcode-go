// 二叉树前序遍历和中序遍历 输出后序遍历

// [1,2,3,4,5,6,7]
// [3,2,4,1,6,5,7]
package main

import (
	"fmt"
)

var mMap map[int]int
var ans []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func visitTree(node *TreeNode) {
	if node == nil {
		return
	}
	visitTree(node.Left)
	visitTree(node.Right)
	ans = append(ans, node.Val)
}
func getTree(pre []int, mid []int, pl int, pr int, ml int, mr int) *TreeNode {
	if pl > pr || ml > mr {
		return nil
	}
	index := mMap[pre[pl]]
	node := &TreeNode{
		Val: pre[pl],
	}
	node.Left = getTree(pre, mid, pl+1, pl+index-ml, ml, index-1)
	node.Right = getTree(pre, mid, pl+index-ml+1, pr, index+1, mr)
	return node
}
func getPostOrder(pre []int, mid []int) []int {
	mMap = make(map[int]int)
	ans = make([]int, 0)
	for index, val := range mid {
		mMap[val] = index
	}
	root := getTree(pre, mid, 0, len(pre)-1, 0, len(pre)-1)
	visitTree(root)
	return ans
}

func main() {
	pre := []int{1, 2, 3, 4, 5, 6, 7}
	mid := []int{3, 2, 4, 1, 6, 5, 7}
	fmt.Println(getPostOrder(pre, mid))
}


kube-proxy
同pod的容器之间 localhost 
容器 bridge
跨物理机的容器 
flannel 隧道 查询路由 
callico 物理机之间建立bgp

