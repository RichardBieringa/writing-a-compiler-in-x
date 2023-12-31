package ast

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
