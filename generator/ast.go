package generator

import (
	"os"
	"path/filepath"
)

const AstTmpl = `
package {{.Package}}

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
`

type AstModel struct {
	Package string
}

func (ctx *Context) GenerateAstGo() error {
	file, err := os.Create(filepath.Join(ctx.Out, "ast.go"))
	if err != nil {
		return err
	}
	defer closeFmt(file)

	executed, _ := execute(AstTmpl, AstModel{Package: ctx.Package})

	_, err = file.WriteString(executed)

	return err
}
