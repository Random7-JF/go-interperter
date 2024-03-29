package ast

import "monkey/token"

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

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode()          {}
func (ls *LetStatement) TokenLiteral() string    { return ls.Token.Literal }
func (i *Identifier) expressionNode()            {}
func (i *Identifier) TokenLiteral() string       { return i.Token.Literal }
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
