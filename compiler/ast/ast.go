package ast

import "github.com/elvodqa/baz/compiler/lexer"

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
	Token lexer.Token // the token.LET token
	Name  *Identifier
	Value Expression
}
