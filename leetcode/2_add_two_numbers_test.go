package leetcode

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewLinkList(vals []int) *ListNode {
	var head, curr, prev *ListNode
	var isHeadSet = false
	for _, val := range vals {
		curr = &ListNode{val, nil}
		if !isHeadSet {
			head = curr
			isHeadSet = true
		}
		if prev != nil {
			prev.Next = curr
		}
		prev = curr
	}
	return head
}

func ListNodeCompare(l, r *ListNode) int {
	for l != nil && r != nil {
		if l.Val > r.Val {
			return 1
		} else if l.Val < r.Val {
			return -1
		}
		l = l.Next
		r = r.Next
	}

	if l == nil && r == nil {
		return 0
	}

	if l == nil {
		return -1
	} else {
		return 1
	}
}

func ToString(l1 *ListNode) string {
	var buffer bytes.Buffer
	var curr *ListNode
	curr = l1

	if l1 == nil {
		return "(nil)"
	}

	buffer.WriteString("(")
	for {
		buffer.WriteString(strconv.Itoa(curr.Val))
		if curr.Next == nil {
			break
		}
		curr = curr.Next
		buffer.WriteString("->")
	}
	buffer.WriteString(")")

	return buffer.String()
}

func TestListNodeCompare(t *testing.T) {
	var golden = []struct {
		want int
		v1   []int
		v2   []int
	}{
		{0, []int{2, 4, 3}, []int{2, 4, 3}},
		{-1, nil, []int{2, 4, 3}},
		{1, []int{2, 4, 3}, nil},
		{-1, []int{2, 4, 3}, []int{5, 6, 4}},
	}

	for _, test := range golden {
		l1 := NewLinkList(test.v1)
		l2 := NewLinkList(test.v2)
		actual := ListNodeCompare(l1, l2)
		assert.Equal(t, test.want, actual, "compare %s and %s but want %d but got %d", ToString(l1), ToString(l2), test.want, actual)
	}
}
func TestAddTwoNumbers(t *testing.T) {
	var golden = []struct {
		name string
		want []int
		v1   []int
		v2   []int
	}{
		{"zero", []int{0}, []int{0}, []int{0}},
		{"normal", []int{7, 9, 7}, []int{2, 3, 3}, []int{5, 6, 4}},
		{"carryinmid", []int{7, 0, 8}, []int{2, 4, 3}, []int{5, 6, 4}},
		{"carryinlast", []int{7, 0, 1, 1}, []int{2, 4, 5}, []int{5, 6, 5}},
		{"l1longer", []int{7, 9, 7, 1}, []int{2, 3, 3, 1}, []int{5, 6, 4}},
		{"l2longer", []int{7, 9, 7, 1}, []int{2, 3, 3}, []int{5, 6, 4, 1}},
		{"l1longerwithcarry", []int{7, 9, 0, 2}, []int{2, 3, 5, 1}, []int{5, 6, 5}},
	}

	for _, test := range golden {
		l1 := NewLinkList(test.v1)
		l2 := NewLinkList(test.v2)
		want := NewLinkList(test.want)
		actual := addTwoNumbers(l1, l2)
		assert.True(t, ListNodeCompare(want, actual) == 0, "%s addTwoNumbers:%v + %v; want %v but got %v", test.name, ToString(l1), ToString(l2), ToString(want), ToString(actual))
	}
}
