// Code generated from GrammarDesc.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // GrammarDesc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGrammarDescListener is a complete listener for a parse tree produced by GrammarDescParser.
type BaseGrammarDescListener struct{}

var _ GrammarDescListener = &BaseGrammarDescListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarDescListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarDescListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarDescListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarDescListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRules is called when production rules is entered.
func (s *BaseGrammarDescListener) EnterRules(ctx *RulesContext) {}

// ExitRules is called when production rules is exited.
func (s *BaseGrammarDescListener) ExitRules(ctx *RulesContext) {}

// EnterLexer_rule is called when production lexer_rule is entered.
func (s *BaseGrammarDescListener) EnterLexer_rule(ctx *Lexer_ruleContext) {}

// ExitLexer_rule is called when production lexer_rule is exited.
func (s *BaseGrammarDescListener) ExitLexer_rule(ctx *Lexer_ruleContext) {}

// EnterSkip_rule is called when production skip_rule is entered.
func (s *BaseGrammarDescListener) EnterSkip_rule(ctx *Skip_ruleContext) {}

// ExitSkip_rule is called when production skip_rule is exited.
func (s *BaseGrammarDescListener) ExitSkip_rule(ctx *Skip_ruleContext) {}

// EnterStart_rule is called when production start_rule is entered.
func (s *BaseGrammarDescListener) EnterStart_rule(ctx *Start_ruleContext) {}

// ExitStart_rule is called when production start_rule is exited.
func (s *BaseGrammarDescListener) ExitStart_rule(ctx *Start_ruleContext) {}

// EnterParser_rule is called when production parser_rule is entered.
func (s *BaseGrammarDescListener) EnterParser_rule(ctx *Parser_ruleContext) {}

// ExitParser_rule is called when production parser_rule is exited.
func (s *BaseGrammarDescListener) ExitParser_rule(ctx *Parser_ruleContext) {}

// EnterTerm_parser_rule_item is called when production term_parser_rule_item is entered.
func (s *BaseGrammarDescListener) EnterTerm_parser_rule_item(ctx *Term_parser_rule_itemContext) {}

// ExitTerm_parser_rule_item is called when production term_parser_rule_item is exited.
func (s *BaseGrammarDescListener) ExitTerm_parser_rule_item(ctx *Term_parser_rule_itemContext) {}

// EnterNon_term_parser_rule_item is called when production non_term_parser_rule_item is entered.
func (s *BaseGrammarDescListener) EnterNon_term_parser_rule_item(ctx *Non_term_parser_rule_itemContext) {
}

// ExitNon_term_parser_rule_item is called when production non_term_parser_rule_item is exited.
func (s *BaseGrammarDescListener) ExitNon_term_parser_rule_item(ctx *Non_term_parser_rule_itemContext) {
}

// EnterCode_item is called when production code_item is entered.
func (s *BaseGrammarDescListener) EnterCode_item(ctx *Code_itemContext) {}

// ExitCode_item is called when production code_item is exited.
func (s *BaseGrammarDescListener) ExitCode_item(ctx *Code_itemContext) {}
