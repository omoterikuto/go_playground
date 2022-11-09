package _2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	ansP := ans
	var up bool

	for l1 != nil || l2 != nil {
		var l1Val int
		if l1 != nil {
			l1Val = l1.Val
		}
		var l2Val int
		if l2 != nil {
			l2Val = l2.Val
		}
		sum := l1Val + l2Val
		if up {
			sum += 1
			up = false
		}
		if sum >= 10 {
			up = true
			sum -= 10
		}
		ansP.Val = sum
		newNode := new(ListNode)
		ansP.Next = newNode
		ansP = newNode
	}

	return ans
}
