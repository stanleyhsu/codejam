package leetcode

import (
	"bytes"
	"math"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val := GetIntegarFromList(l1) + GetIntegarFromList(l2)
	digitArray := ConvertIntToIntArray(val)
	return NewLinkList(digitArray)
}

func GetIntegarFromList(l *ListNode) int {
	var power, result int
	var curr *ListNode
	if l == nil {
		return 0
	}

	curr = l
	for {
		result = result + curr.Val*int(math.Pow10(power))
		if curr.Next == nil {
			break
		}
		curr = curr.Next
		power += 1
	}
	return result
}

func ConvertIntToIntArray(val int) []int {
	var digitArray []int
	var remainder, modulus int

	remainder = val % 10
	modulus = val / 10
	digitArray = append(digitArray, remainder)

	for modulus > 0 {
		val = modulus
		remainder = val % 10
		modulus = val / 10
		digitArray = append(digitArray, remainder)
	}
	return digitArray
}

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

func ToString(l1 *ListNode) string {
	var buffer bytes.Buffer
	var curr *ListNode
	curr = l1

	if l1 == nil {
		return ""
	}
	for {
		buffer.WriteString(strconv.Itoa(curr.Val))
		if curr.Next == nil {
			break
		}
		curr = curr.Next
		buffer.WriteString(" -> ")
	}
	return buffer.String()
}
