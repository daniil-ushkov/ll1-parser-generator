package kotlin

import (
	"github.com/stretchr/testify/assert"
	"ll1-parser-generator/tests/kotlin/parser"
	"testing"
)

func equal(t *testing.T, expected, actual parser.Ast) {
	expectedNonTerminal, okExpectedNonTerminal := expected.(*parser.AstNonTerminal)
	expectedTerminal, okExpectedTerminal := expected.(*parser.AstTerminal)

	actualNonTerminal, okActualNonTerminal := actual.(*parser.AstNonTerminal)
	actualTerminal, okActualTerminal := actual.(*parser.AstTerminal)

	if okExpectedNonTerminal && okActualNonTerminal {
		assert.Equal(t, expectedNonTerminal.NonTerminal, actualNonTerminal.NonTerminal)

		assert.Equal(t, len(expectedNonTerminal.Children), len(actualNonTerminal.Children))

		for i := range expectedNonTerminal.Children {
			equal(t, expectedNonTerminal.Children[i], actualNonTerminal.Children[i])
		}

		return
	}

	if okExpectedTerminal && okActualTerminal {
		assert.Equal(t, expectedTerminal.Token.Symbol, actualTerminal.Token.Symbol)
		assert.Equal(t, expectedTerminal.Token.Lexeme, actualTerminal.Token.Lexeme)

		assert.Equal(t, len(expectedTerminal.Children), len(actualTerminal.Children))

		for i := range expectedTerminal.Children {
			equal(t, expectedTerminal.Children[i], actualTerminal.Children[i])
		}

		return
	}

	assert.Fail(t, "ouuch")
}

func TestEmptyParam(t *testing.T) {
	lexer := parser.NewLexer("fun foo()")
	astParser := parser.Parser{Lexer: lexer}
	actual, err := astParser.Parse()

	assert.Nil(t, err)

	expected := &parser.AstNonTerminal{
		NonTerminal: parser.Sig,
		AstBase: parser.AstBase{
			Children: []parser.Ast{
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.FUN, Lexeme: "fun"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "foo"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.LP, Lexeme: "("}},
				&parser.AstNonTerminal{NonTerminal: parser.Params},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.RP, Lexeme: ")"}},
				&parser.AstNonTerminal{NonTerminal: parser.Ret},
			},
		},
	}

	equal(t, expected, actual)
}

func TestOneParam(t *testing.T) {
	lexer := parser.NewLexer("fun foo(x: Int)")
	astParser := parser.Parser{Lexer: lexer}
	actual, err := astParser.Parse()

	assert.Nil(t, err)

	expected := &parser.AstNonTerminal{
		NonTerminal: parser.Sig,
		AstBase: parser.AstBase{
			Children: []parser.Ast{
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.FUN, Lexeme: "fun"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "foo"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.LP, Lexeme: "("}},
				&parser.AstNonTerminal{NonTerminal: parser.Params, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "x"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
					}}},
					&parser.AstNonTerminal{NonTerminal: parser.Params1},
				}}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.RP, Lexeme: ")"}},
				&parser.AstNonTerminal{NonTerminal: parser.Ret},
			},
		},
	}

	equal(t, expected, actual)
}

func TestRet(t *testing.T) {
	lexer := parser.NewLexer("fun foo(): Int")
	astParser := parser.Parser{Lexer: lexer}
	actual, err := astParser.Parse()

	assert.Nil(t, err)

	expected := &parser.AstNonTerminal{
		NonTerminal: parser.Sig,
		AstBase: parser.AstBase{
			Children: []parser.Ast{
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.FUN, Lexeme: "fun"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "foo"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.LP, Lexeme: "("}},
				&parser.AstNonTerminal{NonTerminal: parser.Params},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.RP, Lexeme: ")"}},
				&parser.AstNonTerminal{NonTerminal: parser.Ret, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
				}}},
			},
		},
	}

	equal(t, expected, actual)
}

func TestOneParamAndRet(t *testing.T) {
	lexer := parser.NewLexer("fun foo(x: Int): Int")
	astParser := parser.Parser{Lexer: lexer}
	actual, err := astParser.Parse()

	assert.Nil(t, err)

	expected := &parser.AstNonTerminal{
		NonTerminal: parser.Sig,
		AstBase: parser.AstBase{
			Children: []parser.Ast{
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.FUN, Lexeme: "fun"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "foo"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.LP, Lexeme: "("}},
				&parser.AstNonTerminal{NonTerminal: parser.Params, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "x"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
					}}},
					&parser.AstNonTerminal{NonTerminal: parser.Params1},
				}}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.RP, Lexeme: ")"}},
				&parser.AstNonTerminal{NonTerminal: parser.Ret, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
				}}},
			},
		},
	}

	equal(t, expected, actual)
}

func TestTwoParamAndRet(t *testing.T) {
	lexer := parser.NewLexer("fun foo(x: Int, y: Int): Int")
	astParser := parser.Parser{Lexer: lexer}
	actual, err := astParser.Parse()

	assert.Nil(t, err)

	expected := &parser.AstNonTerminal{
		NonTerminal: parser.Sig,
		AstBase: parser.AstBase{
			Children: []parser.Ast{
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.FUN, Lexeme: "fun"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "foo"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.LP, Lexeme: "("}},
				&parser.AstNonTerminal{NonTerminal: parser.Params, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "x"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
					}}},
					&parser.AstNonTerminal{NonTerminal: parser.Params1, AstBase: parser.AstBase{Children: []parser.Ast{
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.COM, Lexeme: ","}},
						&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "y"}},
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
						}}},
						&parser.AstNonTerminal{NonTerminal: parser.Params1},
					}}},
				}}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.RP, Lexeme: ")"}},
				&parser.AstNonTerminal{NonTerminal: parser.Ret, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
				}}},
			},
		},
	}

	equal(t, expected, actual)
}

func TestThreeParamAndRet(t *testing.T) {
	lexer := parser.NewLexer("fun foo( x: Int, y: Int, z: Int): Int")
	astParser := parser.Parser{Lexer: lexer}
	actual, err := astParser.Parse()

	assert.Nil(t, err)

	expected := &parser.AstNonTerminal{
		NonTerminal: parser.Sig,
		AstBase: parser.AstBase{
			Children: []parser.Ast{
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.FUN, Lexeme: "fun"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "foo"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.LP, Lexeme: "("}},
				&parser.AstNonTerminal{NonTerminal: parser.Params, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "x"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
					}}},
					&parser.AstNonTerminal{NonTerminal: parser.Params1, AstBase: parser.AstBase{Children: []parser.Ast{
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.COM, Lexeme: ","}},
						&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "y"}},
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
						}}},
						&parser.AstNonTerminal{NonTerminal: parser.Params1, AstBase: parser.AstBase{Children: []parser.Ast{
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.COM, Lexeme: ","}},
							&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
								&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "z"}},
								&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
								&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
							}}},
							&parser.AstNonTerminal{NonTerminal: parser.Params1},
						}}},
					}}},
				}}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.RP, Lexeme: ")"}},
				&parser.AstNonTerminal{NonTerminal: parser.Ret, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
				}}},
			},
		},
	}

	equal(t, expected, actual)
}

func TestThreeParamAndRetWithSpaces(t *testing.T) {
	lexer := parser.NewLexer(`				fun
foo			
			( x					
:	 Int, y			: 
Int		, z			: 
Int)		: 		Int

`)
	astParser := parser.Parser{Lexer: lexer}
	actual, err := astParser.Parse()

	assert.Nil(t, err)

	expected := &parser.AstNonTerminal{
		NonTerminal: parser.Sig,
		AstBase: parser.AstBase{
			Children: []parser.Ast{
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.FUN, Lexeme: "fun"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "foo"}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.LP, Lexeme: "("}},
				&parser.AstNonTerminal{NonTerminal: parser.Params, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "x"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
					}}},
					&parser.AstNonTerminal{NonTerminal: parser.Params1, AstBase: parser.AstBase{Children: []parser.Ast{
						&parser.AstTerminal{Token: parser.Token{Symbol: parser.COM, Lexeme: ","}},
						&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "y"}},
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
						}}},
						&parser.AstNonTerminal{NonTerminal: parser.Params1, AstBase: parser.AstBase{Children: []parser.Ast{
							&parser.AstTerminal{Token: parser.Token{Symbol: parser.COM, Lexeme: ","}},
							&parser.AstNonTerminal{NonTerminal: parser.Param, AstBase: parser.AstBase{Children: []parser.Ast{
								&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "z"}},
								&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
								&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
							}}},
							&parser.AstNonTerminal{NonTerminal: parser.Params1},
						}}},
					}}},
				}}},
				&parser.AstTerminal{Token: parser.Token{Symbol: parser.RP, Lexeme: ")"}},
				&parser.AstNonTerminal{NonTerminal: parser.Ret, AstBase: parser.AstBase{Children: []parser.Ast{
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.COL, Lexeme: ":"}},
					&parser.AstTerminal{Token: parser.Token{Symbol: parser.ID, Lexeme: "Int"}},
				}}},
			},
		},
	}

	equal(t, expected, actual)
}
