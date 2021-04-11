package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%.3f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("%s%c%s", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	var sb strings.Builder
	sb.WriteString(c.fn)
	sb.WriteString("(")
	for i, arg := range c.args {
		sb.WriteString(arg.String())
		if i < len(c.args)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	return sb.String()
}
