package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode() {}

// Token Type which is the TOKEN.LET
// Second is the name of the Identifier type
// Third is the Expression. Ths is given to be evaluated to the evaluator which returns the value of the expression
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Token type
// Value contains the string literal of that identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
