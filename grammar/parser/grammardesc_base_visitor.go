// Code generated from GrammarDesc.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // GrammarDesc

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGrammarDescVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGrammarDescVisitor) VisitRules(ctx *RulesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarDescVisitor) VisitLexer_rule(ctx *Lexer_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarDescVisitor) VisitSkip_rule(ctx *Skip_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarDescVisitor) VisitStart_rule(ctx *Start_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarDescVisitor) VisitParser_rule(ctx *Parser_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarDescVisitor) VisitTerm_parser_rule_item(ctx *Term_parser_rule_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarDescVisitor) VisitNon_term_parser_rule_item(ctx *Non_term_parser_rule_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarDescVisitor) VisitCode_item(ctx *Code_itemContext) interface{} {
	return v.VisitChildren(ctx)
}
