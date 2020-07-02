package palindrome

import (
	"sort"
	"testing"
)

func TestIsPalindrome_Even(t *testing.T) {
	seq := []int{4, 2, 2, 4}
	seqInter := sort.IntSlice(seq)
	res := IsPalindrome(seqInter)
	if !res {
		t.Log("Unexpected value: ", res)
		t.Fail()
	}
}

func TestIsPalindrome_Odd(t *testing.T) {
	seq := []int{4, 2, 3, 2, 4}
	seqInter := sort.IntSlice(seq)
	res := IsPalindrome(seqInter)
	if !res {
		t.Log("Unexpected value: ", res)
		t.Fail()
	}
}
