package palindrome

import (
	"sort"
)

// IsPalindrome takes a sort.Interface and uses its methods to determine if the given sequence is a palindrome
// [a, b, b, c] => true
// "examplestring" => false
func IsPalindrome(s sort.Interface) bool {
	i := 0
	j := s.Len() - 1

	for i < j {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
		i++
		j--
	}
	return true
}
