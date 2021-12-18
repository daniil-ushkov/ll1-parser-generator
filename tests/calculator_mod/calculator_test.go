package calculator_mod

import (
	"github.com/stretchr/testify/assert"
	"ll1-parser-generator/tests/calculator_mod/parser"
	"testing"
)

func Test1(t *testing.T) {
	lexer := parser.NewLexer("8 // 2")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 3, attr["val"].(int))
}

func Test2(t *testing.T) {
	lexer := parser.NewLexer("81 // 8 // 2")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 4, attr["val"].(int))
}

func Test3(t *testing.T) {
	lexer := parser.NewLexer("1 + -3")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, -2, attr["val"].(int))
}

func Test4(t *testing.T) {
	lexer := parser.NewLexer("-2 + -3 + -5 + -4 / 2")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, -12, attr["val"].(int))
}

func Test6(t *testing.T) {
	lexer := parser.NewLexer("81 // 3 * 4 // 2")
	expected := 8
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, expected, attr["val"].(int))
}
