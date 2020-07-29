package eval

import (
	"testing"
)

func TestOpExpr(t *testing.T) {

	// create tree
	x := literal(4)
	y := literal(5)
	binExpr := binary{op: '+', x: x, y: y}

	// give to concrete type
	opexpr := OpExpr{binExpr}
	got := opexpr.Min()

	if got != literal(4) {
		t.Error("Unexpected result: " + got.String())
		t.Fail()
	}
}
