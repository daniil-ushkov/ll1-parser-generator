// Code generated from GrammarDesc.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // GrammarDesc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GrammarDescParser.
type GrammarDescVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarDescParser#rules.
	VisitRules(ctx *RulesContext) interface{}

	// Visit a parse tree produced by GrammarDescParser#lexer_rule.
	VisitLexer_rule(ctx *Lexer_ruleContext) interface{}

	// Visit a parse tree produced by GrammarDescParser#skip_rule.
	VisitSkip_rule(ctx *Skip_ruleContext) interface{}

	// Visit a parse tree produced by GrammarDescParser#start_rule.
	VisitStart_rule(ctx *Start_ruleContext) interface{}

	// Visit a parse tree produced by GrammarDescParser#parser_rule.
	VisitParser_rule(ctx *Parser_ruleContext) interface{}

	// Visit a parse tree produced by GrammarDescParser#term_parser_rule_item.
	VisitTerm_parser_rule_item(ctx *Term_parser_rule_itemContext) interface{}

	// Visit a parse tree produced by GrammarDescParser#non_term_parser_rule_item.
	VisitNon_term_parser_rule_item(ctx *Non_term_parser_rule_itemContext) interface{}

	// Visit a parse tree produced by GrammarDescParser#code_item.
	VisitCode_item(ctx *Code_itemContext) interface{}
}
