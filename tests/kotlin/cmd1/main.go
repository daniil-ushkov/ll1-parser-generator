package main

import (
	"ll1-parser-generator/generator"
	"ll1-parser-generator/grammar"
)

const GrammarDesc = `
$FUN : "fun";
$ID  : "[a-zA-Z][a-zA-Z0-9]*";
$COL : ":";
$COM : ",";
$LP  : "\\(";
$RP  : "\\)";

skip: "[ \t\n]+";

start: $sig;

$sig -> $FUN $ID $LP $params $RP $ret;

$params -> $param $params1;
$params -> ;

$param -> $ID $COL $ID;

$params1 -> $COM $param $params1;
$params1 -> ;

$ret -> $COL $ID;
$ret -> ;
`

func main() {
	desc, _ := grammar.ParseDesc(GrammarDesc)
	gr := grammar.New(desc)
	ctx := generator.Context{
		Package: "parser",
		Out:     "tests/kotlin/parser",
		Grammar: gr,
	}

	_ = ctx.GenerateAstGo()
	_ = ctx.GenerateLexerGo()
	_ = ctx.GenerateAstParserGo()
	_ = ctx.GenerateSymbolsGo()
}
