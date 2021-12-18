package grammar

type LexerRule struct {
	Terminal Terminal
	Regexp   Regexp
}

type Regexp string

type SkipRule struct {
	Regexp Regexp
}

type ParserRule struct {
	NonTerminal NonTerminal
	Rule        []ParserRuleItem
}

func (r *ParserRule) Striped() Rule {
	var rule Rule
	for _, item := range r.Rule {
		switch item1 := item.(type) {
		case TerminalParserRuleItem:
			rule = append(rule, item1.Terminal)
		case NonTerminalParserRuleItem:
			rule = append(rule, item1.NonTerminal)
		}
	}

	return rule
}

type ParserRuleItem interface {
	isParserRuleItem()
}

type TerminalParserRuleItem struct {
	Alias    string
	Terminal Terminal
}

func (t TerminalParserRuleItem) isParserRuleItem() {}

type NonTerminalParserRuleItem struct {
	Alias       string
	NonTerminal NonTerminal
	Attrs       string
}

func (n NonTerminalParserRuleItem) isParserRuleItem() {}

type Code string

func (c Code) isParserRuleItem() {}
