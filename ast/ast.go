package ast

import (
	"fmt"
	"strconv"
)

type ExprNode interface {
	ToString() string
}

type ValueNode struct {
	Value int64
}

func (v *ValueNode) ToString() string {
	return strconv.FormatInt(v.Value, 10)
}

type PrefixOperatorNode struct {
	Operator string
	Rhs      ExprNode
}

func (p *PrefixOperatorNode) ToString() string {
	return fmt.Sprintf("%s%s", p.Operator, p.Rhs.ToString())
}

type InfixOperatorNode struct {
	Lhs      ExprNode
	Operator string
	Rhs      ExprNode
}

func (i *InfixOperatorNode) ToString() string {
	// we treat suffix expression as infix expression, so we must handle a nullable situation
	if i.Rhs != nil {
		return fmt.Sprintf("(%s %s %s)", i.Lhs.ToString(), i.Operator, i.Rhs.ToString())
	} else {
		return fmt.Sprintf("(%s%s)", i.Lhs.ToString(), i.Operator)
	}
}
