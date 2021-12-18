package grammar

type NonTerminalSet map[NonTerminal]struct{}

func NewNonTerminalSet() NonTerminalSet {
	return map[NonTerminal]struct{}{}
}

func (s NonTerminalSet) Add(symbol NonTerminal) {
	s[symbol] = struct{}{}
}

func (s NonTerminalSet) AddAll(set NonTerminalSet) {
	for symbol := range set {
		s.Add(symbol)
	}
}

func (s NonTerminalSet) Erase(symbol NonTerminal) {
	delete(s, symbol)
}

func (s NonTerminalSet) Has(symbol NonTerminal) bool {
	_, ok := s[symbol]
	return ok
}

func (s NonTerminalSet) Len() int {
	return len(s)
}

func (s NonTerminalSet) Empty() bool {
	return len(s) == 0
}
