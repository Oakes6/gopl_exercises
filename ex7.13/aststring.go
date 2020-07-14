package eval

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (lit literal) String() string {
	return fmt.Sprintf("%.6f", lit)
}

func (uni unary) String() string {
	return fmt.Sprintf("%s%s", uni.op, uni.x.String())
}

func (bin binary) String() string {
	return fmt.Sprintf("%s %s %s", bin.x.String(), bin.op, bin.y.String())
}

func (c call) String() string {
	if c.fn == "pow" {
		return fmt.Sprintf("%s(%s,%s)", c.fn, c.args[0], c.args[1])
	}
	return fmt.Sprintf("%s(%s)", c.fn, c.args[0])
}
