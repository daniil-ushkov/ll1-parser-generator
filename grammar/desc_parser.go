package grammar

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"ll1-parser-generator/grammar/parser"
	"strings"
)

func ParseDesc(stringDesc string) (*Desc, error) {
	is := antlr.NewInputStream(stringDesc)
	lexer := parser.NewGrammarDescLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	grammarParser := parser.NewGrammarDescParser(stream)

	visitor := newGrammarDescVisitor()
	visitor.Visit(grammarParser.Rules())

	return visitor.desc, nil
}

type grammarDescVisitor struct {
	*parser.BaseGrammarDescVisitor
	desc    *Desc
	counter int
}

func newGrammarDescVisitor() grammarDescVisitor {
	return grammarDescVisitor{
		BaseGrammarDescVisitor: &parser.BaseGrammarDescVisitor{},
		desc: &Desc{
			LexerRules:  []LexerRule{},
			SkipRules:   []SkipRule{},
			ParserRules: []ParserRule{},
		},
	}
}

//func (v *grammarDescVisitor) Visit(tree antlr.ParseTree) interface{} {
//	switch ctx := tree.(type) {
//	case *parser.RulesContext:
//		v.VisitRules(ctx)
//		return nil
//	case *parser.Lexer_ruleContext:
//		v.VisitLexer_rule(ctx)
//		return nil
//	case *parser.Parser_ruleContext:
//		v.VisitParser_rule(ctx)
//		return nil
//	case *parser.Terminal_symbolContext:
//		v.VisitTerminal_symbol(ctx)
//		return nil
//	case *parser.Non_terminal_symbolContext:
//		v.VisitNon_terminal_symbol(ctx)
//		return nil
//	default:
//		panic("ouuch")
//	}
//}

func (v *grammarDescVisitor) Visit(tree antlr.Tree) interface{} {
	switch ctx := tree.(type) {
	case *parser.RulesContext:
		return v.VisitRules(ctx)
	case *parser.Lexer_ruleContext:
		return v.VisitLexer_rule(ctx)
	case *parser.Skip_ruleContext:
		return v.VisitSkip_rule(ctx)
	case *parser.Start_ruleContext:
		return v.VisitStart_rule(ctx)
	case *parser.Parser_ruleContext:
		return v.VisitParser_rule(ctx)
	case *parser.Term_parser_rule_itemContext:
		return v.VisitTerm_parser_rule_item(ctx)
	case *parser.Non_term_parser_rule_itemContext:
		return v.VisitNon_term_parser_rule_item(ctx)
	case *parser.Code_itemContext:
		return v.VisitCode_item(ctx)
	default:
		panic("unreachable")
	}
}

func (v *grammarDescVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		v.Visit(child)
	}
	return nil
}

func (v *grammarDescVisitor) VisitRules(ctx *parser.RulesContext) interface{} {
	v.VisitChildren(ctx)
	return nil
}

func (v *grammarDescVisitor) VisitLexer_rule(ctx *parser.Lexer_ruleContext) interface{} {
	v.desc.LexerRules = append(v.desc.LexerRules, LexerRule{
		Terminal: Terminal{
			Name: ctx.TERM_ID().GetText()[1:],
		},
		Regexp: Regexp(ctx.REGEXP().GetText()),
	})
	return nil
}

func (v *grammarDescVisitor) VisitSkip_rule(ctx *parser.Skip_ruleContext) interface{} {
	v.desc.SkipRules = append(v.desc.SkipRules, SkipRule{
		Regexp: Regexp(ctx.REGEXP().GetText()),
	})
	return nil
}

func (v *grammarDescVisitor) VisitStart_rule(ctx *parser.Start_ruleContext) interface{} {
	v.desc.StartRule = NonTerminal{Name: extractNonTerminalName(ctx.NON_TERM_ID().GetText())}
	return nil
}

func (v *grammarDescVisitor) VisitParser_rule(ctx *parser.Parser_ruleContext) interface{} {
	nonTerminalName := extractNonTerminalName(ctx.NON_TERM_ID().GetText())
	nonTerminal := NonTerminal{Name: nonTerminalName}

	var rule []ParserRuleItem

	for _, child := range ctx.GetChildren()[2 : len(ctx.GetChildren())-1] {
		rule = append(rule, v.Visit(child).(ParserRuleItem))
	}

	v.desc.ParserRules = append(v.desc.ParserRules, ParserRule{NonTerminal: nonTerminal, Rule: rule})

	return nil
}

func (v *grammarDescVisitor) VisitTerm_parser_rule_item(ctx *parser.Term_parser_rule_itemContext) interface{} {
	var alias string
	if ctx.ID() != nil {
		alias = ctx.ID().GetText()
	} else {
		alias = fmt.Sprintf("__unnamed%d", v.counter)
		v.counter++
	}

	name := extractNonTerminalName(ctx.TERM_ID().GetText())

	return TerminalParserRuleItem{
		Terminal: Terminal{Name: name},
		Alias:    alias,
	}
}

func (v *grammarDescVisitor) VisitNon_term_parser_rule_item(ctx *parser.Non_term_parser_rule_itemContext) interface{} {

	var attrs string

	var alias string
	if ctx.ID() != nil {
		alias = ctx.ID().GetText()
	} else {
		alias = fmt.Sprintf("__unnamed%d", v.counter)
		v.counter++
	}

	name := extractNonTerminalName(ctx.NON_TERM_ID().GetText())

	if ctx.ATTR() != nil {
		attrs = extractAttrsOrCode(ctx.ATTR().GetText())
	}

	return NonTerminalParserRuleItem{
		NonTerminal: NonTerminal{Name: name},
		Alias:       alias,
		Attrs:       attrs,
	}
}

func (v *grammarDescVisitor) VisitCode_item(ctx *parser.Code_itemContext) interface{} {
	return Code(strings.TrimSpace(extractAttrsOrCode(ctx.CODE().GetText())))
}

func extractNonTerminalName(parsedName string) string {
	return strings.Title(parsedName[1:])
}

func extractAttrsOrCode(parsedAttrs string) string {
	return parsedAttrs[1 : len(parsedAttrs)-1]
}
