// Code generated from GrammarDesc.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // GrammarDesc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GrammarDescListener is a complete listener for a parse tree produced by GrammarDescParser.
type GrammarDescListener interface {
	antlr.ParseTreeListener

	// EnterRules is called when entering the rules production.
	EnterRules(c *RulesContext)

	// EnterLexer_rule is called when entering the lexer_rule production.
	EnterLexer_rule(c *Lexer_ruleContext)

	// EnterSkip_rule is called when entering the skip_rule production.
	EnterSkip_rule(c *Skip_ruleContext)

	// EnterStart_rule is called when entering the start_rule production.
	EnterStart_rule(c *Start_ruleContext)

	// EnterParser_rule is called when entering the parser_rule production.
	EnterParser_rule(c *Parser_ruleContext)

	// EnterTerm_parser_rule_item is called when entering the term_parser_rule_item production.
	EnterTerm_parser_rule_item(c *Term_parser_rule_itemContext)

	// EnterNon_term_parser_rule_item is called when entering the non_term_parser_rule_item production.
	EnterNon_term_parser_rule_item(c *Non_term_parser_rule_itemContext)

	// EnterCode_item is called when entering the code_item production.
	EnterCode_item(c *Code_itemContext)

	// ExitRules is called when exiting the rules production.
	ExitRules(c *RulesContext)

	// ExitLexer_rule is called when exiting the lexer_rule production.
	ExitLexer_rule(c *Lexer_ruleContext)

	// ExitSkip_rule is called when exiting the skip_rule production.
	ExitSkip_rule(c *Skip_ruleContext)

	// ExitStart_rule is called when exiting the start_rule production.
	ExitStart_rule(c *Start_ruleContext)

	// ExitParser_rule is called when exiting the parser_rule production.
	ExitParser_rule(c *Parser_ruleContext)

	// ExitTerm_parser_rule_item is called when exiting the term_parser_rule_item production.
	ExitTerm_parser_rule_item(c *Term_parser_rule_itemContext)

	// ExitNon_term_parser_rule_item is called when exiting the non_term_parser_rule_item production.
	ExitNon_term_parser_rule_item(c *Non_term_parser_rule_itemContext)

	// ExitCode_item is called when exiting the code_item production.
	ExitCode_item(c *Code_itemContext)
}
