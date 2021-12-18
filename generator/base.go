package generator

import (
	"ll1-parser-generator/grammar"
	"os"
)

type Context struct {
	Out     string
	Package string
	Grammar *grammar.Grammar
	file    *os.File
}
