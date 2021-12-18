package calculator

import (
	"github.com/stretchr/testify/assert"
	"ll1-parser-generator/tests/calculator/parser"
	"testing"
)

func TestSimplePlus(t *testing.T) {
	lexer := parser.NewLexer("1 + 2")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 3, attr["val"].(int))
}

func TestSimpleMul(t *testing.T) {
	lexer := parser.NewLexer("2 * 3")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 6, attr["val"].(int))
}

func TestMulAndPlus1(t *testing.T) {
	lexer := parser.NewLexer("2 * (3 + 4)")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 14, attr["val"].(int))
}

func TestMulAndPlus2(t *testing.T) {
	lexer := parser.NewLexer("2 * (3 + 4) * 5 + 6")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 76, attr["val"].(int))
}

func TestMulAndPlus3(t *testing.T) {
	lexer := parser.NewLexer("1 * 2 * 3 * 4 * 5")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 120, attr["val"].(int))
}

func TestMulAndPlus4(t *testing.T) {
	lexer := parser.NewLexer("(((1 * 2) * 3) * 4) * 5")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 120, attr["val"].(int))
}

func TestMulAndPlus5(t *testing.T) {
	lexer := parser.NewLexer("1 + 2 + 3 + 4 + 5")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 15, attr["val"].(int))
}

func TestMulAndPlus6(t *testing.T) {
	lexer := parser.NewLexer("(((1 + 2) + 3) + 4) + 5")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 15, attr["val"].(int))
}

func TestMulAndPlus7(t *testing.T) {
	lexer := parser.NewLexer("(((((1 + 2) + 3) * ((1 + 2) + 3)) * ((1 + 2) + 3)) * ((1 + 2) + 3)) * ((1 + 2) + 3)")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 7776, attr["val"].(int))
}

func TestMulAndPlus8(t *testing.T) {
	lexer := parser.NewLexer("11 * (5 * 2 + 3 * 21) + 74 * 45 + 123 * (123 + 43) + 1 * 2 * 3 * 4 * 5 * (1 + 2 + 3 + 4 + 5) + (((1 * 2) * 3) * 4) * 5")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 26471, attr["val"].(int))
}

func TestSimpleDiv(t *testing.T) {
	lexer := parser.NewLexer("8 / 2")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 4, attr["val"].(int))
}

func TestDiv(t *testing.T) {
	lexer := parser.NewLexer("120 / 5 / 4 / 3 / 2")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 1, attr["val"].(int))
}

func TestMulAndDiv1(t *testing.T) {
	lexer := parser.NewLexer("8 / 2 * 6 / 3 * 11")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 88, attr["val"].(int))
}

func TestAll1(t *testing.T) {
	lexer := parser.NewLexer("(8 / 2 / 2 + 11) * (1 - 2 - 3) / 2")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, -26, attr["val"].(int))
}

func TestSimpleSub(t *testing.T) {
	lexer := parser.NewLexer("8 - 2")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 6, attr["val"].(int))
}

func TestSub(t *testing.T) {
	lexer := parser.NewLexer("1 - 2 - 3")
	attrParser := parser.AttrParser{Lexer: lexer}
	attr, err := attrParser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, -4, attr["val"].(int))
}
