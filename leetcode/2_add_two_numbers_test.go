package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewLinkList(t *testing.T) {
	var golden = []struct {
		want  string
		input []int
	}{
		{"2 -> 4 -> 3", []int{2, 4, 3}},
		{"5 -> 6 -> 4", []int{5, 6, 4}},
	}

	for _, test := range golden {
		if got := ToString(NewLinkList(test.input)); strings.Compare(got, test.want) != 0 {
			t.Errorf("NewLinkList(%v) Got(%s), Expected(%s)", test.input, got, test.want)
		}
	}
}

func TestGetIntegarFromList(t *testing.T) {
	var golden = []struct {
		want  int
		input []int
	}{
		{342, []int{2, 4, 3}},
		{465, []int{5, 6, 4}},
	}
	for _, test := range golden {
		if got := GetIntegarFromList(NewLinkList(test.input)); got != test.want {
			t.Errorf("GetIntegarFromList(%v) Got(%d), Expected(%d)", test.input, got, test.want)
		}
	}
}

func TestConvertIntToIntArray(t *testing.T) {
	var golden = []struct {
		want  []int
		input int
	}{
		{[]int{2, 4, 3}, 342},
		{[]int{5, 6, 4}, 465},
		//{[]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, 1000000000000000000000466},
	}
	for _, test := range golden {
		if got := ConvertIntToIntArray(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("ConvertIntToIntArray(%v) Got(%d), Expected(%d)", test.input, got, test.want)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	var golden = []struct {
		want string
		v1   []int
		v2   []int
	}{
		{"7 -> 0 -> 8", []int{2, 4, 3}, []int{5, 6, 4}},
		//{"7 -> 0 -> 8", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, []int{5, 6, 4}},
	}

	for _, test := range golden {
		l1 := NewLinkList(test.v1)
		l2 := NewLinkList(test.v2)
		if got := ToString(addTwoNumbers(l1, l2)); strings.Compare(got, test.want) != 0 {
			t.Errorf("addTwoNumbers(%v, %v) Got(%s), Expected(%s)", test.v1, test.v2, got, test.want)
		}
	}
}
