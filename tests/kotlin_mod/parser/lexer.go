// Code generated by daniil-ushkov LL(1) parser generator; DO NOT EDIT.
package parser

import "regexp"

type Token struct {
	Symbol Terminal
	Lexeme string
}

type Rule struct {
	Symbol Terminal
	Regexp *regexp.Regexp
}

type Lexer struct {
	Rules []Rule
	Skip  []*regexp.Regexp

	Cur  *Token
	Next *Token

	Input string
}

func NewLexer(input string) Lexer {
	FUNreg, _ := regexp.Compile("fun")
	IDreg, _ := regexp.Compile("[a-zA-Z][a-zA-Z0-9]*")
	COLreg, _ := regexp.Compile(":")
	COMreg, _ := regexp.Compile(",")
	LPreg, _ := regexp.Compile("\\(")
	RPreg, _ := regexp.Compile("\\)")
	skip0, _ := regexp.Compile("[ \t\n]+")

	lex := Lexer{
		Rules: []Rule{
			{Symbol: FUN, Regexp: FUNreg},
			{Symbol: ID, Regexp: IDreg},
			{Symbol: COL, Regexp: COLreg},
			{Symbol: COM, Regexp: COMreg},
			{Symbol: LP, Regexp: LPreg},
			{Symbol: RP, Regexp: RPreg},
		},
		Skip: []*regexp.Regexp{
			skip0,
		},
		Input: input,
	}

	lex.Cur = lex.nextImpl()
	lex.Next = lex.nextImpl()

	return lex
}

func (l *Lexer) HasNext() bool {
	return !(l.Cur.Symbol == EOF && l.Next.Symbol == EOF)
}

func (l *Lexer) MoveNext() {
	l.Cur = l.Next
	l.Next = l.nextImpl()
}

func (l *Lexer) nextImpl() *Token {
	for {
		noWs := true
		for _, skip := range l.Skip {
			inds := skip.FindStringIndex(l.Input)
			if inds != nil && inds[0] == 0 {
				l.Input = l.Input[inds[1]:]
				noWs = false
			}
		}
		if noWs {
			break
		}
	}

	if l.Input == "" {
		return &Token{Symbol: EOF, Lexeme: ""}
	}

	for _, rule := range l.Rules {
		inds := rule.Regexp.FindStringIndex(l.Input)
		if inds != nil && inds[0] == 0 {
			res := &Token{Symbol: rule.Symbol, Lexeme: l.Input[inds[0]:inds[1]]}
			l.Input = l.Input[inds[1]:]
			return res
		}
	}

	return nil
}

func (l *Lexer) Get() Token {
	return *l.Cur
}