package grammar

type TerminalSet map[Terminal]struct{}

func NewTerminalSet() TerminalSet {
	return map[Terminal]struct{}{}
}

func (s TerminalSet) Add(symbol Terminal) {
	s[symbol] = struct{}{}
}

func (s TerminalSet) AddAll(set TerminalSet) {
	for symbol := range set {
		s.Add(symbol)
	}
}

func (s TerminalSet) Erase(symbol Terminal) {
	delete(s, symbol)
}

func (s TerminalSet) Has(symbol Terminal) bool {
	_, ok := s[symbol]
	return ok
}

func (s TerminalSet) Len() int {
	return len(s)
}

func (s TerminalSet) Empty() bool {
	return len(s) == 0
}
