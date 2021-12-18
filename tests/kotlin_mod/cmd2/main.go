package main

import (
	"github.com/goccy/go-graphviz"
	"ll1-parser-generator/tests/kotlin_mod/parser"
	"log"
)

func main() {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}

	lexer := parser.NewLexer("fun foo(x: Int, y: Double, z: Stream): Int")
	astParser := parser.Parser{Lexer: lexer}
	ast, err := astParser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	_, err = ast.AsGraph(graph)

	if err := g.RenderFilename(graph, graphviz.PNG, "tests/kotlin_mod/cmd2/tree.png"); err != nil {
		log.Fatal(err)
	}
}
