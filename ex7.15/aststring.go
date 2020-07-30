package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (lit literal) String() string {
	return fmt.Sprintf("%g", lit)
}

func (uni unary) String() string {
	return fmt.Sprintf("%c%s", uni.op, uni.x.String())
}

func (bin binary) String() string {
	return fmt.Sprintf("(%s %c %s)", bin.x.String(), bin.op, bin.y.String())
}

func (c call) String() string {
	strBuilder := &strings.Builder{}
	strBuilder.WriteString(c.fn)
	strBuilder.WriteString("(")
	for i, v := range c.args {
		if i != 0 {
			strBuilder.WriteString(", ")
		}
		strBuilder.WriteString(v.String())
	}
	strBuilder.WriteString(")")
	return strBuilder.String()
}
