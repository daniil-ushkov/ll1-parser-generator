package grammar

type Symbol interface {
	isSymbol()
}

type Terminal struct {
	Name string
}

func (t Terminal) isSymbol() {}

type NonTerminal struct {
	Name string
}

func (n NonTerminal) isSymbol() {}
