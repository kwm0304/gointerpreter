package ast

import "monkey/token"

//every node has to implement the Node interface(has to provide TokenLiteral)
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

//will be root node of every AST produced
//method tracks 3 fields: var name(identifier), expression that produes the value, and token
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct{
	Token token.Token // the token.LET token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal}

type Identifier struct {
	Token token.Token //the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()		{}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal}