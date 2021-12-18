package parser

import (
	"fmt"
	"github.com/goccy/go-graphviz/cgraph"
)

type Ast interface {
	isAst()
	AddChild(Ast)
	AsGraph(graph *cgraph.Graph) (*cgraph.Node, error)
}

type AstBase struct {
	Children []Ast
}

func (a *AstBase) isAst() {}

func (a *AstBase) AddChild(ast Ast) {
	a.Children = append(a.Children, ast)
}

type AstNonTerminal struct {
	AstBase
	NonTerminal NonTerminal
}

var counter = 0

func (n *AstNonTerminal) AsGraph(graph *cgraph.Graph) (*cgraph.Node, error) {
	node, err := graph.CreateNode(fmt.Sprint(counter))
	counter++
	if err != nil {
		return nil, err
	}

	node.SetLabel(n.NonTerminal.String())

	for _, child := range n.Children {
		childNode, err := child.AsGraph(graph)
		if err != nil {
			return nil, err
		}

		_, err = graph.CreateEdge(fmt.Sprint(counter), node, childNode)
		counter++
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}

type AstTerminal struct {
	AstBase
	Token Token
}

func (n *AstTerminal) AsGraph(graph *cgraph.Graph) (*cgraph.Node, error) {
	node, err := graph.CreateNode(fmt.Sprint(counter))
	counter++
	if err != nil {
		return nil, err
	}

	node.SetLabel(n.Token.Symbol.String() + " : " + n.Token.Lexeme)
	for _, child := range n.Children {
		childNode, err := child.AsGraph(graph)
		if err != nil {
			return nil, err
		}

		_, err = graph.CreateEdge(fmt.Sprint(counter), node, childNode)
		counter++
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}
