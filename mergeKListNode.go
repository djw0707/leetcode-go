
import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}
type PriorityQueue []*ListNode

func (h PriorityQueue) Less(i int, j int) bool {
	return h[i].Val < h[j].Val
}
func (h PriorityQueue) Swap(i int, j int) { h[i], h[j] = h[j], h[i] }
func (h PriorityQueue) Len() int          { return len(h) }
func (h *PriorityQueue) Push(node any)    { *h = append(*h, node.(*ListNode)) }

// node的类型可以使用any或者interface
func (h *PriorityQueue) Pop() any {
	// 数组最后一个才是最小值
	if len(*h) < 1 {
		return nil
	}
	first := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return first
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy_ptr := &ListNode{}
	cur := dummy_ptr
	h := make(PriorityQueue, 0)
	// 初始化需要使用指针
	heap.Init(&h)
	for _, node := range lists {
		if node != nil {
			heap.Push(&h, node)
		}
	}
	for h.Len() > 0 {
		top := heap.Pop(&h).(*ListNode)
		if top.Next != nil {
			heap.Push(&h, top.Next)
		}
		cur.Next = top
		cur = cur.Next
	}
	return dummy_ptr.Next
}