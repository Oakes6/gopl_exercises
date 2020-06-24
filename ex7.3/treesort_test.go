package treesort

import (
	"testing"
)

func TestTreeString(t *testing.T) {
	var root *tree = &tree{value: 4, left: &tree{value: 6}, right: &tree{value: 3}}

	actual := root.String()
	expected := "6->4->3->"
	if actual != expected {
		t.Log("UNEXPECTED RESULT: ", actual)
		t.Fail()
	}
}
