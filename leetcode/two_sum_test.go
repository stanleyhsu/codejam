package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expectIndex := []int{0, 1}

	gotIndex := twoSum(nums, target)

	if !Equal(gotIndex, expectIndex) {
		t.Errorf("TwoSum(%v, %v) GotIndex(%v) Expected(%v)", nums, target, gotIndex, expectIndex)
	}

}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expectIndex := []int{1, 2}

	gotIndex := twoSum(nums, target)

	if !Equal(gotIndex, expectIndex) {
		t.Errorf("TwoSum(%v, %v) GotIndex(%v) Expected(%v)", nums, target, gotIndex, expectIndex)
	}

}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
