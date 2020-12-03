package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(prev *ListNode, newVal int) (curr *ListNode, carrier int) {
	if newVal >= 10 {
		carrier = 1
		newVal -= 10
	} else {
		carrier = 0
	}
	curr = &ListNode{newVal, nil}

	if prev != nil {
		prev.Next = curr
	}

	return curr, carrier
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, curr, prev, left *ListNode
	var carrier, newVal int

	for l1 != nil && l2 != nil {
		newVal = l1.Val + l2.Val + carrier

		curr, carrier = addNode(prev, newVal)
		if head == nil {
			head = curr
		}
		prev = curr
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		left = l1
	} else {
		left = l2
	}
	for left != nil {
		newVal = left.Val + carrier
		curr, carrier = addNode(prev, newVal)
		prev = curr
		left = left.Next
	}

	if carrier == 1 {
		addNode(prev, carrier)
	}
	return head
}
