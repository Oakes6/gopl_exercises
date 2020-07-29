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

	// create tree for 'pow(x,3) + pow(y,2)'
	varX := Var("x")
	x = literal(2)
	fn := "pow"
	args := []Expr{varX, x}
	callObjectX := call{fn: fn, args: args}

	varY := Var("y")
	y = literal(3)
	args = []Expr{varY, y}
	callObjectY := call{fn: fn, args: args}

	op := '+'
	binaryObject := binary{op: op, x: callObjectX, y: callObjectY}

	// assert
	opexpr = OpExpr{binaryObject}
	got = opexpr.Min()

	if got != literal(2) {
		t.Error("Unexpected result: " + got.String())
		t.Fail()
	}

}
