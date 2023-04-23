package ast

import (
	"Monkey/token"
	"bytes"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement
type LetStatement struct {
	Token token.Token // token.LET 词法单元
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (l LetStatement) statementNode() {}

func (l LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

// Identifier
type Identifier struct {
	Token token.Token // token.IDENT 词法单元
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i Identifier) expressionNode() {}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.TokenLiteral() + " ")

	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (r *ReturnStatement) statementNode() {}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

// ExpressionStatement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}

	return ""
}

func (e *ExpressionStatement) statementNode() {}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

// IntegerLiteral
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) expressionNode() {}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

// PrefixExpression
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")

	return out.String()
}

func (p *PrefixExpression) expressionNode() {}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

// InfixExpression
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")

	return out.String()
}

func (i *InfixExpression) expressionNode() {}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

// Boolean
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// IfExpression
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.Consequence.String())

	if i.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(i.Alternative.String())
	}

	return out.String()
}

func (i *IfExpression) expressionNode() {}

func (i *IfExpression) TokenLiteral() string {
	return i.Token.Literal
}

// BlockStatement
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range b.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (b *BlockStatement) statementNode() {}

func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())

	return out.String()
}

func (f *FunctionLiteral) expressionNode() {
	//TODO implement me
	panic("implement me")
}

// CallExpression
type CallExpression struct {
	Token     token.Token // '('词法单元
	Function  Expression  // 标识符或函数字面量
	Arguments []Expression
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
