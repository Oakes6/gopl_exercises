package eval

import (
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "sqrt((A / pi))"},
		{"pow(x, 4)", Env{"x": 4}, "pow(x, 4)"},
		{"5 / 4 * (F - 32)", Env{"F": 100}, "((5 / 4) * (F - 32))"},
	}

	for _, v := range tests {
		expr, err := Parse(v.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := expr.String()
		if got != v.want {
			t.Error("INVALID VALUE\tGOT: " + got + " EXPECTED: " + v.want)
			continue
		}
	}
}
