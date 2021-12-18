package grammar

type Desc struct {
	LexerRules  []LexerRule
	SkipRules   []SkipRule
	StartRule   NonTerminal
	ParserRules []ParserRule
}

func (desc *Desc) Terminals() TerminalSet {
	res := NewTerminalSet()
	for _, rule := range desc.LexerRules {
		res.Add(rule.Terminal)
	}

	return res
}

func (desc *Desc) NonTerminals() NonTerminalSet {
	res := NewNonTerminalSet()
	for _, rule := range desc.ParserRules {
		res.Add(rule.NonTerminal)
	}

	return res
}

type Rule []Symbol

var (
	EPS = Terminal{Name: "EPS"}
	EOF = Terminal{Name: "EOF"}
)

type Grammar struct {
	Desc               *Desc
	RulesByNonTerminal map[NonTerminal][]Rule
	First              map[NonTerminal]TerminalSet
	Follow             map[NonTerminal]TerminalSet
}

func New(desc *Desc) *Grammar {
	g := &Grammar{Desc: desc}

	g.RulesByNonTerminal = g.buildRulesByNonTerminal()
	g.First = g.buildFirst()
	g.Follow = g.buildFollow()

	return g
}

func (g *Grammar) buildRulesByNonTerminal() map[NonTerminal][]Rule {
	res := map[NonTerminal][]Rule{}
	for _, parserRule := range g.Desc.ParserRules {
		if _, ok := res[parserRule.NonTerminal]; !ok {
			res[parserRule.NonTerminal] = []Rule{}
		}

		res[parserRule.NonTerminal] = append(res[parserRule.NonTerminal], parserRule.Striped())
	}

	return res
}

func (g *Grammar) buildFirst() map[NonTerminal]TerminalSet {
	first := map[NonTerminal]TerminalSet{}

	for {
		stop := true
		for nonTerminal, rules := range g.RulesByNonTerminal {
			for _, rule := range rules {
				if len(rule) == 0 {
					first[nonTerminal].Add(EPS)
					continue
				}
				switch symbol := rule[0].(type) {
				case Terminal:
					if _, ok := first[nonTerminal]; !ok {
						first[nonTerminal] = NewTerminalSet()
					}

					if !first[nonTerminal].Has(symbol) {
						first[nonTerminal].Add(symbol)
						stop = false
					}
				case NonTerminal:
					if _, ok := first[nonTerminal]; !ok {
						first[nonTerminal] = NewTerminalSet()
					}

					oldLen := first[nonTerminal].Len()
					first[nonTerminal].AddAll(first[symbol])
					newLen := first[nonTerminal].Len()

					if oldLen != newLen {
						stop = false
					}
				}

			}
		}

		if stop {
			break
		}
	}

	return first
}

func (g *Grammar) GetFirst(rule Rule) TerminalSet {
	s := NewTerminalSet()

	for _, symbol := range rule {
		if terminal, ok := symbol.(Terminal); ok {
			s.Add(terminal)
			return s
		}

		if nonTerminal, ok := symbol.(NonTerminal); ok {
			s.AddAll(g.First[nonTerminal])
			if !s.Empty() {
				return s
			}
		}
	}

	s.Add(EPS)

	return s
}

func (g *Grammar) buildFollow() map[NonTerminal]TerminalSet {
	follow := map[NonTerminal]TerminalSet{}

	nonTerminals := g.Desc.NonTerminals()
	for nonTerminal := range nonTerminals {
		follow[nonTerminal] = NewTerminalSet()
	}

	follow[g.Desc.StartRule].Add(EOF)

	changed := true
	for changed {
		changed = false
		for _, rule := range g.Desc.ParserRules {
			ruleStriped := rule.Striped()
			for i, symbol := range ruleStriped {
				if nonTerminal, ok := symbol.(NonTerminal); ok {
					oldLen := follow[nonTerminal].Len()

					firstAfter := g.GetFirst(ruleStriped[i+1:])
					hasEps := firstAfter.Has(EPS)
					firstAfter.Erase(EPS)

					follow[nonTerminal].AddAll(firstAfter)

					if hasEps {
						followNonTerminal := follow[rule.NonTerminal]
						follow[nonTerminal].AddAll(followNonTerminal)
					}

					if oldLen != follow[nonTerminal].Len() {
						changed = true
					}
				}
			}
		}
	}

	return follow
}
